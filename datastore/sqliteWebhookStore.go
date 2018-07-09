package datastore

import "github.com/swagchat/chat-api/model"

func (p *sqliteProvider) createWebhookStore() {
	rdbCreateWebhookStore(p.database)
}

func (p *sqliteProvider) SelectWebhooks(event model.WebhookEventType, opts ...WebhookOption) ([]*model.Webhook, error) {
	return rdbSelectWebhooks(p.database, event, opts...)
}
