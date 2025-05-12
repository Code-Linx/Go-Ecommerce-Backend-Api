package main

import (
	"database/sql"
	"log"

	"github.com/Code-Linx/Go-Ecommerce-Backend-Api/cmd/api"
	"github.com/Code-Linx/Go-Ecommerce-Backend-Api/config"
	"github.com/Code-Linx/Go-Ecommerce-Backend-Api/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}
	InitStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Sucessfully Connceted!")
}
