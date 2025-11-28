package api

import (
	"chat/internal/handler/user"
	logger "chat/internal/util"
	"encoding/json"
	"net/http"
)

func CreateRegisterRequest(w http.ResponseWriter, r *http.Request) {
	createRequest(w, r, user.Register, "register user")
}
func CreateLoginRequest(w http.ResponseWriter, r *http.Request) {
	createRequest(w, r, user.Login, "log user in")
}

func createRequest[v any](w http.ResponseWriter, r *http.Request,
	f func(http.ResponseWriter, *http.Request, v), operation string) {
	var req v
	if err := decodeJSON(w, r, &req); err != nil {
		logger.Info("Failed to " + operation)
		return
	} else {
		f(w, r, req)
	}
}

func decodeJSON(w http.ResponseWriter, r *http.Request, v any) error {
	var err error
	if err = json.NewDecoder(r.Body).Decode(&v); err != nil {
		logger.Info("Failed to decode JSON")
		logger.Debug(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	return err
}
