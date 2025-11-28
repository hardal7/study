package user

import (
	"chat/internal/model"
	"chat/internal/repository"
	logger "chat/internal/util"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost int = 10
)

func Register(w http.ResponseWriter, r *http.Request, rr model.RegisterRequest) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(rr.Password), bcryptCost)
	if err != nil {
		logger.Info("Failed to create user: Could not hash password")
		logger.Debug(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := model.User{
		Email:     rr.Email,
		Username:  rr.Username,
		Password:  string(passwordHash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	isDuplicate, err := repository.IsDuplicateUser(r.Context(), user)
	if err != nil {
		logger.Info("Failed to check if user " + user.Username + " is duplicate")
		logger.Debug(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else if isDuplicate {
		logger.Info("User " + user.Username + " is already registered")
		http.Error(w, "User is already registered", http.StatusBadRequest)
	} else {
		if err := repository.CreateUser(r.Context(), user); err != nil {
			logger.Info("Failed to create user: " + rr.Username)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			logger.Info("Created user: " + rr.Username)
			w.WriteHeader(http.StatusCreated)
		}
	}
}
