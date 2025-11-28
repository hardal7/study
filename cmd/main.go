package main

import (
	"chat/internal/api"
	"chat/internal/config"
	"chat/internal/repository"
	logger "chat/internal/util"
)

func init() {
	logger.Init()
	config.Load()
}

func main() {
	repository.CreateDBConnection()
	api.RunAPIServer()
}
