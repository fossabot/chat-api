package rest

import (
	"net/http"
	"net/url"

	"github.com/go-zoo/bone"
	"github.com/swagchat/chat-api/model"
	"github.com/swagchat/chat-api/service"
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
	var req model.CreateRoomRequest
	if err := decodeBody(r, &req); err != nil {
		respondJSONDecodeError(w, r, "")
		return
	}

	room, errRes := service.CreateRoom(r.Context(), &req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusCreated, "application/json", room)
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	req := &model.GetRoomsRequest{}
	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		respondErr(w, r, http.StatusBadRequest, nil)
		return
	}

	limit, offset, orders, errRes := setPagingParams(params)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	req.Limit = limit
	req.Offset = offset
	req.Orders = orders

	rooms, errRes := service.GetRooms(r.Context(), req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", rooms)
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	req := &model.GetRoomRequest{}

	roomID := bone.GetValue(r, "roomId")
	req.RoomID = roomID

	room, errRes := service.GetRoom(r.Context(), req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", room)
}

func putRoom(w http.ResponseWriter, r *http.Request) {
	var req model.UpdateRoomRequest
	if err := decodeBody(r, &req); err != nil {
		respondJSONDecodeError(w, r, "")
		return
	}

	req.RoomID = bone.GetValue(r, "roomId")

	room, errRes := service.UpdateRoom(r.Context(), &req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", room)
}

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	req := &model.DeleteRoomRequest{}

	roomID := bone.GetValue(r, "roomId")
	req.RoomID = roomID

	errRes := service.DeleteRoom(r.Context(), req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusNoContent, "", nil)
}

func getRoomMessages(w http.ResponseWriter, r *http.Request) {
	req := &model.GetRoomMessagesRequest{}

	roomID := bone.GetValue(r, "roomId")
	req.RoomID = roomID

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		respondErr(w, r, http.StatusBadRequest, nil)
		return
	}

	limit, offset, orders, errRes := setPagingParams(params)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	req.Limit = limit
	req.Offset = offset
	req.Orders = orders

	messages, errRes := service.GetRoomMessages(r.Context(), req)
	if errRes != nil {
		respondError(w, r, errRes)
		return
	}

	respond(w, r, http.StatusOK, "application/json", messages)
}
