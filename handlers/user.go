package handlers

import (
	"net/http"

	"log"

	"github.com/fairway-corp/swagchat-api/models"
	"github.com/fairway-corp/swagchat-api/services"
	"github.com/go-zoo/bone"
)

func SetUserMux() {
	Mux.PostFunc("/users", colsHandler(aclHandler(PostUser)))
	Mux.GetFunc("/users", colsHandler(GetUsers))
	Mux.GetFunc("/users/#userId^[a-z0-9-]$", colsHandler(GetUser))
	Mux.PutFunc("/users/#userId^[a-z0-9-]$", colsHandler(PutUser))
	Mux.DeleteFunc("/users/#userId^[a-z0-9-]$", colsHandler(DeleteUser))
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Context().Value("role"))
	var post models.User
	if err := decodeBody(r, &post); err != nil {
		respondJsonDecodeError(w, r, "Create user item")
		return
	}

	user, pd := services.PostUser(&post)
	if pd != nil {
		respondErr(w, r, pd.Status, pd)
		return
	}

	respond(w, r, http.StatusCreated, "application/json", user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, pd := services.GetUsers()
	if pd != nil {
		respondErr(w, r, pd.Status, pd)
		return
	}

	respond(w, r, http.StatusOK, "application/json", users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := bone.GetValue(r, "userId")
	user, pd := services.GetUser(userId)
	if pd != nil {
		respondErr(w, r, pd.Status, pd)
		return
	}

	respond(w, r, http.StatusOK, "application/json", user)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	var put models.User
	if err := decodeBody(r, &put); err != nil {
		respondJsonDecodeError(w, r, "Update user item")
		return
	}

	put.UserId = bone.GetValue(r, "userId")
	user, pd := services.PutUser(&put)
	if pd != nil {
		respondErr(w, r, pd.Status, pd)
		return
	}

	respond(w, r, http.StatusOK, "application/json", user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := bone.GetValue(r, "userId")
	pd := services.DeleteUser(userId)
	if pd != nil {
		respondErr(w, r, pd.Status, pd)
		return
	}

	respond(w, r, http.StatusNoContent, "", nil)
}
