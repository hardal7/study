package main

import (
	"github.com/hardal7/study/internal/api"
	"github.com/hardal7/study/internal/config"
	"github.com/hardal7/study/internal/repository"
	logger "github.com/hardal7/study/internal/util"
)

func init() {
	logger.Init()
	config.Load()
}

func main() {
	repository.CreateDBConnection()
	api.RunAPIServer()
}
