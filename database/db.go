package database

import (
	"database/sql"
	"fmt"

	"github.com/aklbnelly/libraryproject/config"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB(envs *config.Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envs.DBHost, envs.DBPort, envs.DBUser, envs.DBPassword, envs.DBName, envs.DBSSLMode)
	fmt.Println("Trying to connect with db") //  Выводим строку подключения

	var err error
	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error opening DB:", err)
		return err
	}
	if err := Db.Ping(); err != nil {
		fmt.Println("Error pinging DB:", err)
		Db.Close()
		return err

	}
	fmt.Println("Connected to DB successfully!")
	return nil
}
