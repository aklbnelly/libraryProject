package main

import (
	"github.com/aklbnelly/libraryproject/config"
	"github.com/aklbnelly/libraryproject/database"
	"github.com/aklbnelly/libraryproject/server"
	"github.com/aklbnelly/libraryproject/utils"
)

func main() {

	projectConfig, err := config.LoadConfig()
	if err != nil {
		utils.Logger.Errorf("problems with envs, error: %v", err)
		return
	}
	utils.InitLoggers(projectConfig.LogFormat)

	err = database.InitDB(projectConfig)
	if err != nil {
		utils.Logger.Error("can not connect to db")
		return
	}

	defer func() {
		if database.Db != nil {
			database.Db.Close()
		}
	}()

	server.Run()
}
