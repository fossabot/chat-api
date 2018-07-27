package datastore

import "github.com/swagchat/chat-api/model"

func (p *gcpSQLProvider) createDeviceStore() {
	rdbCreateDeviceStore(p.ctx, p.database)
}

func (p *gcpSQLProvider) InsertDevice(device *model.Device) (*model.Device, error) {
	return rdbInsertDevice(p.ctx, p.database, device)
}

func (p *gcpSQLProvider) SelectDevices(userID string) ([]*model.Device, error) {
	return rdbSelectDevices(p.ctx, p.database, userID)
}

func (p *gcpSQLProvider) SelectDevice(userID string, platform int32) (*model.Device, error) {
	return rdbSelectDevice(p.ctx, p.database, userID, platform)
}

func (p *gcpSQLProvider) SelectDevicesByUserID(userID string) ([]*model.Device, error) {
	return rdbSelectDevicesByUserID(p.ctx, p.database, userID)
}

func (p *gcpSQLProvider) SelectDevicesByToken(token string) ([]*model.Device, error) {
	return rdbSelectDevicesByToken(p.ctx, p.database, token)
}

func (p *gcpSQLProvider) UpdateDevice(device *model.Device) error {
	return rdbUpdateDevice(p.ctx, p.database, device)
}

func (p *gcpSQLProvider) DeleteDevice(userID string, platform int32) error {
	return rdbDeleteDevice(p.ctx, p.database, userID, platform)
}
