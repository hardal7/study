package api

import (
	"chat/internal/config"
	"chat/internal/middleware"
	logger "chat/internal/util"
	"net/http"
)

func RunAPIServer() {
	router := http.NewServeMux()
	router.HandleFunc("POST /register", middleware.LogRequest(CreateRegisterRequest))
	router.HandleFunc("POST /login", middleware.LogRequest(CreateLoginRequest))

	logger.Info("Starting server on port: " + config.App.Port)
	server := http.Server{
		Addr:    ":" + config.App.Port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Failed to start server")
		logger.Debug(err.Error())
	}
}
