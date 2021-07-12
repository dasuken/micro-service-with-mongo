package resthandlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io"
	"microservices/api/restutil"
	"microservices/pb"
	"net/http"
	"time"
)

type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type authHandlers struct {
	authSrvClient pb.AuthServiceClient
}

func NewAuthHandlers(authSrvClient pb.AuthServiceClient) AuthHandlers {
	return &authHandlers{authSrvClient:authSrvClient}
}

func (h *authHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := new(pb.User)
	if err := json.Unmarshal(body, user); err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user.Created = time.Now().Unix()
	user.Updated = user.Created
	user.Id = bson.NewObjectId().Hex()

	resp, err := h.authSrvClient.SignUp(context.Background(), user)
	if err := json.Unmarshal(body, user); err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	restutil.WriteAsJson(w, http.StatusOK, resp)
}

func (h *authHandlers) PutUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := new(pb.User)
	if err := json.Unmarshal(body, user); err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	vars := mux.Vars(r)
	user.Id = vars["id"]
	resp, err := h.authSrvClient.UpdateUser(context.Background(), user)
	if err := json.Unmarshal(body, user); err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	restutil.WriteAsJson(w, http.StatusOK, resp)
}

func (h *authHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resp, err := h.authSrvClient.GetUser(context.Background(), &pb.GetUserRequest{Id: vars["id"]})
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	restutil.WriteAsJson(w, http.StatusOK, resp)
}


func (h *authHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	stream, err := h.authSrvClient.ListUsers(context.Background(), &pb.ListUsersRequest{})
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var users []*pb.User

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
		}
		users = append(users, user)
	}

	restutil.WriteAsJson(w, http.StatusOK, users)
}

func (h *authHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resp, err := h.authSrvClient.DeleteUser(context.Background(), &pb.GetUserRequest{Id:vars["id"]})
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}

	restutil.WriteAsJson(w, http.StatusOK, resp)
}