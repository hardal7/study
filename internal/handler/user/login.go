package user

import (
	"chat/internal/config"
	"chat/internal/model"
	"chat/internal/repository"
	logger "chat/internal/util"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	jwtExpirationDays int = 30
)

func Login(w http.ResponseWriter, r *http.Request, lr model.LoginRequest) {
	user, err := repository.GetUser(r.Context(), lr)
	if err != nil {
		logger.Info("Failed to fetch user")
		logger.Debug(err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lr.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		logger.Info("Login attempt with incorect password")
		http.Error(w, "Incorrect Password", http.StatusUnauthorized)
		return
	} else if err != nil {
		logger.Info("Failed to compare password to hash")
		logger.Debug(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * time.Duration(jwtExpirationDays)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.App.JWT_SECRET))
	if err != nil {
		logger.Info("Failed to generate JWT token.")
		logger.Debug(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
		if err != nil {
			logger.Info("Failed to generate JSON response")
			logger.Debug(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			logger.Info("Logged user and sent token")
		}
	}
}
