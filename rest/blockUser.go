package rest

import (
	"net/http"

	"github.com/go-zoo/bone"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/swagchat/chat-api/model"
	"github.com/swagchat/chat-api/service"
)

func setBlockUserMux() {
	mux.GetFunc("/users/#userId^[a-z0-9-]$/blocks", commonHandler(selfResourceAuthzHandler(getBlockUsers)))
	mux.PutFunc("/users/#userId^[a-z0-9-]$/blocks", commonHandler(selfResourceAuthzHandler(putBlockUsers)))
	mux.DeleteFunc("/users/#userId^[a-z0-9-]$/blocks", commonHandler(selfResourceAuthzHandler(deleteBlockUsers)))
}

func getBlockUsers(w http.ResponseWriter, r *http.Request) {
	span, _ := opentracing.StartSpanFromContext(r.Context(), "rest.getBlockUsers")
	defer span.Finish()

	userID := bone.GetValue(r, "userId")

	blockUsers, pd := service.GetBlockUsers(r.Context(), userID)
	if pd != nil {
		respondErr(w, r, pd.Status, pd)
		return
	}

	respond(w, r, http.StatusOK, "application/json", blockUsers)
}

func putBlockUsers(w http.ResponseWriter, r *http.Request) {
	span, _ := opentracing.StartSpanFromContext(r.Context(), "rest.putBlockUsers")
	defer span.Finish()

	var reqUIDs model.RequestBlockUserIDs
	if err := decodeBody(r, &reqUIDs); err != nil {
		respondJSONDecodeError(w, r, "")
		return
	}

	userID := bone.GetValue(r, "userId")

	blockUsers, pd := service.PutBlockUsers(r.Context(), userID, &reqUIDs)
	if pd != nil {
		respondErr(w, r, pd.Status, pd)
		return
	}

	respond(w, r, http.StatusOK, "application/json", blockUsers)
}

func deleteBlockUsers(w http.ResponseWriter, r *http.Request) {
	span, _ := opentracing.StartSpanFromContext(r.Context(), "rest.deleteBlockUsers")
	defer span.Finish()

	var reqUIDs model.RequestBlockUserIDs
	if err := decodeBody(r, &reqUIDs); err != nil {
		respondJSONDecodeError(w, r, "")
		return
	}

	userID := bone.GetValue(r, "userId")

	blockUsers, pd := service.DeleteBlockUsers(r.Context(), userID, &reqUIDs)
	if pd != nil {
		respondErr(w, r, pd.Status, pd)
		return
	}

	respond(w, r, http.StatusOK, "application/json", blockUsers)
}
