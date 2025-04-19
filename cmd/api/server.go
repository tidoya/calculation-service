package main

import (
	"fmt"
	"net/http"
)

func StartServer(port string, handler http.Handler) error {
	fmt.Printf("Сервер запущен на http://localhost%s\n", port)
	return http.ListenAndServe(port, handler)
}

		// Вызываем функцию из server.go
//     	err := StartServer(":8080")
//     	if err != nil {
//     		panic(err) // или обработайте ошибку по-другому
//     	}