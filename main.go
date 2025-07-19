package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Настройка обработчиков
	http.HandleFunc("/api/hello", helloHandler)

	// Обслуживание статических файлов из папки static
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Получение порта из переменной окружения (для VK Cloud)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию для локальной разработки
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Печатаем в консоль сервера
	fmt.Println("Hello World!")

	// Отправляем ответ клиенту (VK Mini App требует JSON)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"response": "Hello World!"}`))
}
