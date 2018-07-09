package datastore

import "github.com/swagchat/chat-api/model"

func (p *sqliteProvider) createSubscriptionStore() {
	rdbCreateSubscriptionStore(p.database)
}

func (p *sqliteProvider) InsertSubscription(room *model.Subscription) (*model.Subscription, error) {
	return rdbInsertSubscription(p.database, room)
}

func (p *sqliteProvider) SelectSubscription(roomID, userID string, platform int) (*model.Subscription, error) {
	return rdbSelectSubscription(p.database, roomID, userID, platform)
}

func (p *sqliteProvider) SelectDeletedSubscriptionsByRoomID(roomID string) ([]*model.Subscription, error) {
	return rdbSelectDeletedSubscriptionsByRoomID(p.database, roomID)
}

func (p *sqliteProvider) SelectDeletedSubscriptionsByUserID(userID string) ([]*model.Subscription, error) {
	return rdbSelectDeletedSubscriptionsByUserID(p.database, userID)
}

func (p *sqliteProvider) SelectDeletedSubscriptionsByUserIDAndPlatform(userID string, platform int) ([]*model.Subscription, error) {
	return rdbSelectDeletedSubscriptionsByUserIDAndPlatform(p.database, userID, platform)
}

func (p *sqliteProvider) DeleteSubscription(subscription *model.Subscription) error {
	return rdbDeleteSubscription(p.database, subscription)
}
