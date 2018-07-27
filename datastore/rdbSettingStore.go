package datastore

import (
	"context"
	"fmt"
	"strconv"

	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/swagchat/chat-api/logger"
	"github.com/swagchat/chat-api/model"
)

func rdbCreateSettingStore(ctx context.Context, db string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "datastore.rdbCreateSettingStore")
	defer span.Finish()

	master := RdbStore(db).master()

	tableMap := master.AddTableWithName(model.Setting{}, tableNameSetting)
	tableMap.SetKeys(true, "id")
	if err := master.CreateTablesIfNotExists(); err != nil {
		logger.Error(fmt.Sprintf("An error occurred while creating setting table. %v.", err))
		return
	}
}

func rdbSelectLatestSetting(ctx context.Context, db string) (*model.Setting, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "datastore.rdbSelectLatestSetting")
	defer span.Finish()

	replica := RdbStore(db).replica()

	var settings []*model.Setting
	nowTimestamp := time.Now().Unix()
	nowTimestampString := strconv.FormatInt(nowTimestamp, 10)
	query := fmt.Sprintf("SELECT * FROM %s WHERE expired=0 OR expired>%s ORDER BY created DESC LIMIT 1;", tableNameSetting, nowTimestampString)
	if _, err := replica.Select(&settings, query); err != nil {
		logger.Error(fmt.Sprintf("An error occurred while getting setting. %v.", err))
		return nil, err
	}
	if len(settings) > 0 {
		return settings[0], nil
	}

	return nil, nil
}
