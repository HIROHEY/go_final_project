package server

import (
	"net/http"

	"github.com/HIROHEY/go_final_project/pkg/serverAndConfig/config"
)

// SetupAndRun настраивает и запускает сервер
func SetupAndRun(port string) error {
	http.Handle("/", http.FileServer(http.Dir(config.WebDir)))
	return http.ListenAndServe(":"+port, nil)

}
