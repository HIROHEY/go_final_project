package api

import (
	"net/http"
)

// Обработчик для /api/task, который определяет действие в зависимости от HTTP-метода
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r) // Добавление новой задачи
	case http.MethodGet:
		tasksHandler(w, r) // Проверка задач

	default:
		http.Error(w, "Указанный метод не найден", http.StatusMethodNotAllowed)
	}
}
