package datastore

import "github.com/swagchat/chat-api/model"

func (p *gcpSQLProvider) createAssetStore() {
	rdbCreateAssetStore(p.ctx, p.database)
}

func (p *gcpSQLProvider) InsertAsset(asset *model.Asset) error {
	return rdbInsertAsset(p.ctx, p.database, asset)
}

func (p *gcpSQLProvider) SelectAsset(assetID string) (*model.Asset, error) {
	return rdbSelectAsset(p.ctx, p.database, assetID)
}
