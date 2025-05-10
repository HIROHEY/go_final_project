package server

import (
	"net/http"

	"github.com/HIROHEY/go_final_project/pkg/serverAndConfing/confing"
)

// SetupAndRun настраивает и запускает сервер
func SetupAndRun(port string) error {
	http.Handle("/", http.FileServer(http.Dir(confing.WebDir)))
	return http.ListenAndServe(":"+port, nil)

}
