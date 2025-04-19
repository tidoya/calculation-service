package main

import (
	"encoding/json"
	"net/http"
)

// Данные, которые будем отдавать в JSON
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
   println("Starting server on :8080")  // Добавьте эту строку


	// Обработчик для маршрута "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовок Content-Type
		w.Header().Set("Content-Type", "application/json")

		// Создаем JSON-ответ
		response := Response{
			Message: "Hello, G2221111O 2LANG!",
			Status:  200,
		}

		// Кодируем структуру в JSON и отправляем
		json.NewEncoder(w).Encode(response)
	})

	// Запускаем сервер на порту 8080
	http.ListenAndServe(":8080", nil)
}