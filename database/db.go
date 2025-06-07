package database

import (
	"database/sql"
	"fmt"

	"github.com/aklbnelly/libraryproject/config"
	"github.com/aklbnelly/libraryproject/utils"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB(envs *config.Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envs.DBHost, envs.DBPort, envs.DBUser, envs.DBPassword, envs.DBName, envs.DBSSLMode)
	utils.Logger.Info("Trying to connect with db") //  Выводим строку подключения

	var err error
	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		utils.Logger.Errorf("Error opening DB: %v", err)
		return err
	}
	if err := Db.Ping(); err != nil {
		utils.Logger.Errorf("Error pinging DB: %v", err)
		Db.Close()
		return err

	}
	utils.Logger.Info("Connected to DB successfully!")
	return nil
}
