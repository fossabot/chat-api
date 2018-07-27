package datastore

import "github.com/swagchat/chat-api/model"

func (p *sqliteProvider) createBlockUserStore() {
	rdbCreateBlockUserStore(p.ctx, p.database)
}

func (p *sqliteProvider) InsertBlockUsers(blockUsers []*model.BlockUser) error {
	return rdbInsertBlockUsers(p.ctx, p.database, blockUsers)
}

func (p *sqliteProvider) SelectBlockUsers(userID string) ([]string, error) {
	return rdbSelectBlockUsers(p.ctx, p.database, userID)
}

func (p *sqliteProvider) SelectBlockUser(userID, blockUserID string) (*model.BlockUser, error) {
	return rdbSelectBlockUser(p.ctx, p.database, userID, blockUserID)
}

func (p *sqliteProvider) DeleteBlockUsers(userID string, blockUserIDs []string) error {
	return rdbDeleteBlockUsers(p.ctx, p.database, userID, blockUserIDs)
}
