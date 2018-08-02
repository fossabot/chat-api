package datastore

import (
	"github.com/pkg/errors"
	"github.com/swagchat/chat-api/logger"
	"github.com/swagchat/chat-api/model"
	scpb "github.com/swagchat/protobuf/protoc-gen-go"
)

func (p *gcpSQLProvider) createDeviceStore() {
	master := RdbStore(p.database).master()
	rdbCreateDeviceStore(p.ctx, master)
}

func (p *gcpSQLProvider) InsertDevice(device *model.Device) error {
	master := RdbStore(p.database).master()
	return rdbInsertDevice(p.ctx, master, device)
}

func (p *gcpSQLProvider) SelectDevices(opts ...SelectDevicesOption) ([]*model.Device, error) {
	replica := RdbStore(p.database).replica()
	return rdbSelectDevices(p.ctx, replica, opts...)
}

func (p *gcpSQLProvider) SelectDevice(userID string, platform scpb.Platform) (*model.Device, error) {
	replica := RdbStore(p.database).replica()
	return rdbSelectDevice(p.ctx, replica, userID, platform)
}

func (p *gcpSQLProvider) SelectDevicesByUserID(userID string) ([]*model.Device, error) {
	replica := RdbStore(p.database).replica()
	return rdbSelectDevicesByUserID(p.ctx, replica, userID)
}

func (p *gcpSQLProvider) SelectDevicesByToken(token string) ([]*model.Device, error) {
	replica := RdbStore(p.database).replica()
	return rdbSelectDevicesByToken(p.ctx, replica, token)
}

func (p *gcpSQLProvider) UpdateDevice(device *model.Device) error {
	master := RdbStore(p.database).master()
	tx, err := master.Begin()
	if err != nil {
		err = errors.Wrap(err, "An error occurred while updating device")
		logger.Error(err.Error())
		return err
	}

	err = rdbUpdateDevice(p.ctx, master, tx, device)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		err = errors.Wrap(err, "An error occurred while updating device")
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (p *gcpSQLProvider) DeleteDevice(userID string, platform scpb.Platform) error {
	master := RdbStore(p.database).master()
	tx, err := master.Begin()
	if err != nil {
		err = errors.Wrap(err, "An error occurred while deleting device")
		logger.Error(err.Error())
		return err
	}

	err = rdbDeleteDevice(p.ctx, master, tx, userID, platform)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		err = errors.Wrap(err, "An error occurred while deleting device")
		logger.Error(err.Error())
		return err
	}

	return nil
}
