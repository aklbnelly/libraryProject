package main

import (
	"github.com/aklbnelly/libraryproject/config"
	"github.com/aklbnelly/libraryproject/database"
	"github.com/aklbnelly/libraryproject/server"
	"github.com/aklbnelly/libraryproject/utils"
)

func main() {
	utils.InitLoggers()
	ProjectConfig, err := config.LoadConfig()
	if err != nil {
		utils.ErrorLog.Printf("problems with envs,error: %v", err)
		return
	}

	err = database.InitDB(ProjectConfig)
	if err != nil {
		utils.ErrorLog.Print("can not connect to db")
		return
	}

	defer func() {
		if database.Db != nil {
			database.Db.Close()
		}
	}()

	server.Run()
}
