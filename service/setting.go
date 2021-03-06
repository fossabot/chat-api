// Business Logic

package service

import (
	"context"
	"net/http"

	"github.com/swagchat/chat-api/datastore"
	"github.com/swagchat/chat-api/model"
	"github.com/betchi/tracer"
)

// GetSetting gets setting
func GetSetting(ctx context.Context) (*model.Setting, *model.ErrorResponse) {
	span := tracer.StartSpan(ctx, "GetSetting", "service")
	defer tracer.Finish(span)

	setting, err := datastore.Provider(ctx).SelectLatestSetting()
	if err != nil {
		return nil, model.NewErrorResponse("Failed to get setting.", http.StatusInternalServerError, model.WithError(err))
	}
	if setting == nil {
		return nil, model.NewErrorResponse("Resource not found.", http.StatusNotFound)
	}

	return setting, nil
}
