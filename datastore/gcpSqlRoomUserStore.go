package datastore

import (
	"github.com/pkg/errors"
	"github.com/swagchat/chat-api/logger"
	"github.com/swagchat/chat-api/model"
)

func (p *gcpSQLProvider) createRoomUserStore() {
	master := RdbStore(p.database).master()
	rdbCreateRoomUserStore(p.ctx, master)
}

func (p *gcpSQLProvider) InsertRoomUsers(roomUsers []*model.RoomUser, opts ...InsertRoomUsersOption) error {
	master := RdbStore(p.database).master()
	tx, err := master.Begin()
	if err != nil {
		err = errors.Wrap(err, "An error occurred while inserting user roles")
		logger.Error(err.Error())
		return err
	}

	err = rdbInsertRoomUsers(p.ctx, master, tx, roomUsers, opts...)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		err = errors.Wrap(err, "An error occurred while inserting user roles")
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (p *gcpSQLProvider) SelectRoomUsers(opts ...SelectRoomUsersOption) ([]*model.RoomUser, error) {
	replica := RdbStore(p.database).replica()
	return rdbSelectRoomUsers(p.ctx, replica, opts...)
}

func (p *gcpSQLProvider) SelectRoomUser(roomID, userID string) (*model.RoomUser, error) {
	replica := RdbStore(p.database).replica()
	return rdbSelectRoomUser(p.ctx, replica, roomID, userID)
}

func (p *gcpSQLProvider) SelectRoomUserOfOneOnOne(myUserID, opponentUserID string) (*model.RoomUser, error) {
	replica := RdbStore(p.database).replica()
	return rdbSelectRoomUserOfOneOnOne(p.ctx, replica, myUserID, opponentUserID)
}

func (p *gcpSQLProvider) SelectUserIDsOfRoomUser(roomID string, opts ...SelectUserIDsOfRoomUserOption) ([]string, error) {
	replica := RdbStore(p.database).replica()
	return rdbSelectUserIDsOfRoomUser(p.ctx, replica, roomID, opts...)
}

func (p *gcpSQLProvider) UpdateRoomUser(roomUser *model.RoomUser) error {
	master := RdbStore(p.database).master()
	return rdbUpdateRoomUser(p.ctx, master, roomUser)
}

func (p *gcpSQLProvider) DeleteRoomUsers(roomID string, userIDs []string) error {
	master := RdbStore(p.database).master()
	tx, err := master.Begin()
	if err != nil {
		err = errors.Wrap(err, "An error occurred while inserting user roles")
		logger.Error(err.Error())
		return err
	}

	err = rdbDeleteRoomUsers(p.ctx, master, tx, roomID, userIDs)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		err = errors.Wrap(err, "An error occurred while inserting user roles")
		logger.Error(err.Error())
		return err
	}

	return nil
}
