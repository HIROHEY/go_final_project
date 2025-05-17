package main

import (
	"log"
	"os"

	"github.com/HIROHEY/go_final_project/pkg/api"
	"github.com/HIROHEY/go_final_project/pkg/db"
	"github.com/HIROHEY/go_final_project/pkg/serverAndConfing/confing"
	"github.com/HIROHEY/go_final_project/pkg/serverAndConfing/server"
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

	port := confing.GetPort()
	if err := server.SetupAndRun(port); err != nil {
		log.Fatalf("Ошибка подключения сервера: %v", err)
	}

	err = server.SetupAndRun(port)

	if err != nil {
		log.Fatal(err)
	}

}
