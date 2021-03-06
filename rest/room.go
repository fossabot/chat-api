package rest

import (
	"net/http"
	"net/url"

	"github.com/go-zoo/bone"
	"github.com/swagchat/chat-api/model"
	"github.com/swagchat/chat-api/service"
	"github.com/betchi/tracer"
)

func setRoomMux() {
	mux.PostFunc("/rooms", commonHandler(postRoom))
	mux.GetFunc("/rooms", commonHandler(adminAuthzHandler(getRooms)))
	mux.GetFunc("/rooms/#roomId^[a-z0-9-]$", commonHandler(roomMemberAuthzHandler(getRoom)))
	mux.PutFunc("/rooms/#roomId^[a-z0-9-]$", commonHandler(roomMemberAuthzHandler(putRoom)))
	mux.DeleteFunc("/rooms/#roomId^[a-z0-9-]$", commonHandler(roomMemberAuthzHandler(deleteRoom)))
	mux.GetFunc("/rooms/#roomId^[a-z0-9-]$/messages", commonHandler(roomMemberAuthzHandler(updateLastAccessedHandler(getRoomMessages))))
}

func postRoom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.StartSpan(ctx, "postRoom", "rest")
	defer tracer.Finish(span)

	var req model.CreateRoomRequest
	if err := decodeBody(r, &req); err != nil {
		respondJSONDecodeError(w, r, "")
		return
	}

	room, errRes := service.CreateRoom(ctx, &req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusCreated, "application/json", room)
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.StartSpan(ctx, "getRooms", "rest")
	defer tracer.Finish(span)

	req := &model.RetrieveRoomsRequest{}
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		errRes := model.NewErrorResponse("", http.StatusBadRequest, model.WithError(err))
		respondError(w, r, errRes)
		return
	}

	limit, offset, _, _, orders, errRes := setPagingParams(params)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	req.Limit = limit
	req.Offset = offset
	req.Orders = orders

	rooms, errRes := service.RetrieveRooms(ctx, req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", rooms)
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.StartSpan(ctx, "getRoom", "rest")
	defer tracer.Finish(span)

	req := &model.RetrieveRoomRequest{}

	roomID := bone.GetValue(r, "roomId")
	req.RoomID = roomID

	room, errRes := service.RetrieveRoom(ctx, req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", room)
}

func putRoom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.StartSpan(ctx, "putRoom", "rest")
	defer tracer.Finish(span)

	var req model.UpdateRoomRequest
	if err := decodeBody(r, &req); err != nil {
		respondJSONDecodeError(w, r, "")
		return
	}

	req.RoomID = bone.GetValue(r, "roomId")

	room, errRes := service.UpdateRoom(ctx, &req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", room)
}

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.StartSpan(ctx, "deleteRoom", "rest")
	defer tracer.Finish(span)

	req := &model.DeleteRoomRequest{}

	roomID := bone.GetValue(r, "roomId")
	req.RoomID = roomID

	errRes := service.DeleteRoom(ctx, req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusNoContent, "", nil)
}

func getRoomMessages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := tracer.StartSpan(ctx, "getRoomMessages", "rest")
	defer tracer.Finish(span)

	req := &model.RetrieveRoomMessagesRequest{}

	roomID := bone.GetValue(r, "roomId")
	req.RoomID = roomID

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		errRes := model.NewErrorResponse("", http.StatusBadRequest, model.WithError(err))
		respondError(w, r, errRes)
		return
	}

	limit, offset, limitTimestamp, offsetTimestamp, orders, errRes := setPagingParams(params)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	req.Limit = limit
	req.Offset = offset
	req.Orders = orders
	req.LimitTimestamp = limitTimestamp
	req.OffsetTimestamp = offsetTimestamp

	messages, errRes := service.RetrieveRoomMessages(ctx, req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", messages)
}
