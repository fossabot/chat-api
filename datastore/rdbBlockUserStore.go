package datastore

import (
	"context"
	"fmt"

	logger "github.com/betchi/zapper"
	"github.com/pkg/errors"
	"github.com/swagchat/chat-api/model"
	"github.com/betchi/tracer"
	gorp "gopkg.in/gorp.v2"
)

func rdbCreateBlockUserStore(ctx context.Context, dbMap *gorp.DbMap) {
	span := tracer.StartSpan(ctx, "rdbCreateBlockUserStore", "datastore")
	defer tracer.Finish(span)

	tableMap := dbMap.AddTableWithName(model.BlockUser{}, tableNameBlockUser)
	tableMap.SetUniqueTogether("user_id", "block_user_id")
	err := dbMap.CreateTablesIfNotExists()
	if err != nil {
		err = errors.Wrap(err, "An error occurred while creating block user table")
		logger.Error(err.Error())
		tracer.SetError(span, err)
		return
	}
}

func rdbInsertBlockUsers(ctx context.Context, dbMap *gorp.DbMap, tx *gorp.Transaction, bus []*model.BlockUser, opts ...InsertBlockUsersOption) error {
	span := tracer.StartSpan(ctx, "rdbInsertBlockUsers", "datastore")
	defer tracer.Finish(span)

	if len(bus) == 0 {
		return nil
	}

	opt := insertBlockUsersOptions{}
	for _, o := range opts {
		o(&opt)
	}

	if opt.beforeClean {
		err := rdbDeleteBlockUsers(ctx, dbMap, tx, DeleteBlockUsersOptionFilterByUserIDs([]string{bus[0].UserID}))
		if err != nil {
			return err
		}
	}

	for _, bu := range bus {
		if !opt.beforeClean {
			existBlockUser, err := rdbSelectBlockUser(ctx, dbMap, bu.UserID, bu.BlockUserID)
			if err != nil {
				return err
			}
			if existBlockUser != nil {
				continue
			}
		}

		err := tx.Insert(bu)
		if err != nil {
			err = errors.Wrap(err, "An error occurred while inserting block users")
			logger.Error(err.Error())
			tracer.SetError(span, err)
			return err
		}
	}

	return nil
}

func rdbSelectBlockUsers(ctx context.Context, dbMap *gorp.DbMap, userID string) ([]*model.MiniUser, error) {
	span := tracer.StartSpan(ctx, "rdbSelectBlockUsers", "datastore")
	defer tracer.Finish(span)

	var blockUsers []*model.MiniUser
	query := fmt.Sprintf(`SELECT
	u.user_id,
	u.name,
	u.picture_url,
	u.information_url,
	u.meta_data,
	u.can_block,
	u.last_accessed,
	u.created,
	u.modified
	FROM %s AS bu 
	LEFT JOIN %s AS u ON bu.block_user_id = u.user_id
	WHERE bu.user_id=:userId;`, tableNameBlockUser, tableNameUser)
	params := map[string]interface{}{
		"userId": userID,
	}
	_, err := dbMap.Select(&blockUsers, query, params)
	if err != nil {
		err = errors.Wrap(err, "An error occurred while getting block users")
		logger.Error(err.Error())
		tracer.SetError(span, err)
		return nil, err
	}

	return blockUsers, nil
}

func rdbSelectBlockUserIDs(ctx context.Context, dbMap *gorp.DbMap, userID string) ([]string, error) {
	span := tracer.StartSpan(ctx, "rdbSelectBlockUserIDs", "datastore")
	defer tracer.Finish(span)

	var blockUserIDs []string
	query := fmt.Sprintf("SELECT block_user_id FROM %s WHERE user_id=:userId;", tableNameBlockUser)
	params := map[string]interface{}{
		"userId": userID,
	}
	_, err := dbMap.Select(&blockUserIDs, query, params)
	if err != nil {
		err = errors.Wrap(err, "An error occurred while getting block userIds")
		logger.Error(err.Error())
		tracer.SetError(span, err)
		return nil, err
	}

	return blockUserIDs, nil
}

func rdbSelectBlockedUsers(ctx context.Context, dbMap *gorp.DbMap, userID string) ([]*model.MiniUser, error) {
	span := tracer.StartSpan(ctx, "rdbSelectBlockedUsers", "datastore")
	defer tracer.Finish(span)

	var blockedUsers []*model.MiniUser
	query := fmt.Sprintf(`SELECT
	u.user_id,
	u.name,
	u.picture_url,
	u.information_url,
	u.meta_data,
	u.can_block,
	u.last_accessed,
	u.created,
	u.modified
	FROM %s AS bu 
	LEFT JOIN %s AS u ON bu.user_id = u.user_id
	WHERE bu.block_user_id=:userId;`, tableNameBlockUser, tableNameUser)
	params := map[string]interface{}{
		"userId": userID,
	}
	_, err := dbMap.Select(&blockedUsers, query, params)
	if err != nil {
		err = errors.Wrap(err, "An error occurred while getting blocked users")
		logger.Error(err.Error())
		tracer.SetError(span, err)
		return nil, err
	}

	return blockedUsers, nil
}

func rdbSelectBlockedUserIDs(ctx context.Context, dbMap *gorp.DbMap, userID string) ([]string, error) {
	span := tracer.StartSpan(ctx, "rdbSelectBlockedUserIDs", "datastore")
	defer tracer.Finish(span)

	var blockUserIDs []string
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE block_user_id=:userId;", tableNameBlockUser)
	params := map[string]interface{}{
		"userId": userID,
	}
	_, err := dbMap.Select(&blockUserIDs, query, params)
	if err != nil {
		err = errors.Wrap(err, "An error occurred while getting blocked userIds")
		logger.Error(err.Error())
		tracer.SetError(span, err)
		return nil, err
	}

	return blockUserIDs, nil
}

func rdbSelectBlockUser(ctx context.Context, dbMap *gorp.DbMap, userID, blockUserID string) (*model.BlockUser, error) {
	span := tracer.StartSpan(ctx, "rdbSelectBlockUser", "datastore")
	defer tracer.Finish(span)

	var blockUsers []*model.BlockUser
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=:userId AND block_user_id=:blockUserId;", tableNameBlockUser)
	params := map[string]interface{}{
		"userId":      userID,
		"blockUserId": blockUserID,
	}
	_, err := dbMap.Select(&blockUsers, query, params)
	if err != nil {
		err = errors.Wrap(err, "An error occurred while getting block user")
		logger.Error(err.Error())
		tracer.SetError(span, err)
		return nil, err
	}

	if len(blockUsers) == 1 {
		return blockUsers[0], nil
	}

	return nil, nil
}

func rdbDeleteBlockUsers(ctx context.Context, dbMap *gorp.DbMap, tx *gorp.Transaction, opts ...DeleteBlockUsersOption) error {
	span := tracer.StartSpan(ctx, "rdbDeleteBlockUsers", "datastore")
	defer tracer.Finish(span)

	opt := deleteBlockUsersOptions{}
	for _, o := range opts {
		o(&opt)
	}

	if len(opt.userIDs) == 0 && len(opt.blockUserIDs) == 0 {
		err := errors.New("An error occurred while deleting block users. Be sure to specify either userIDs or blockUserIDs")
		logger.Error(err.Error())
		tracer.SetError(span, err)
		return err
	}

	if len(opt.userIDs) > 0 {
		userIDsQuery, userIDsParams := makePrepareExpressionForInOperand(opt.userIDs)
		query := fmt.Sprintf("DELETE FROM %s WHERE user_id IN (%s)", tableNameBlockUser, userIDsQuery)
		_, err := tx.Exec(query, userIDsParams...)
		if err != nil {
			err = errors.Wrap(err, "An error occurred while deleting block users")
			logger.Error(err.Error())
			tracer.SetError(span, err)
			return err
		}
	}

	if len(opt.blockUserIDs) > 0 {
		blockUserIDsQuery, blockUserIDsParams := makePrepareExpressionForInOperand(opt.blockUserIDs)
		query := fmt.Sprintf("DELETE FROM %s WHERE block_user_id IN (%s)", tableNameBlockUser, blockUserIDsQuery)
		_, err := tx.Exec(query, blockUserIDsParams...)
		if err != nil {
			err = errors.Wrap(err, "An error occurred while deleting block users")
			logger.Error(err.Error())
			tracer.SetError(span, err)
			return err
		}
	}

	return nil
}
