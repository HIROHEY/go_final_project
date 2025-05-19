package main

import (
	"log"
	"os"

	"github.com/HIROHEY/go_final_project/pkg/api"
	"github.com/HIROHEY/go_final_project/pkg/db"
	"github.com/HIROHEY/go_final_project/pkg/serverAndConfig/config"
	"github.com/HIROHEY/go_final_project/pkg/serverAndConfig/server"
)

func main() {
	api.Init()

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	err := db.Init(dbFile)
	if err != nil {
		log.Fatalf("Ошибка инициализации базы %v", err)
	}

	port := config.GetPort()
	if err := server.SetupAndRun(port); err != nil {
		log.Fatalf("Ошибка подключения сервера: %v", err)
	}

	if err != nil {
		log.Fatal(err)
	}

}
