package api

import (
	"net/http"
)

// Обработчик для /api/task, который определяет действие в зависимости от HTTP-метода
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodGet:
		// Если параметр id отсутствует — возвращаем ошибку
		if !r.URL.Query().Has("id") {
			writeJson(w, map[string]string{"error": "параметр id обязателен"}, http.StatusBadRequest)
			return
		}
		getTaskHandler(w, r) // Только для запросов с id
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
