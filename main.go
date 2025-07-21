package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	// Создаем мультиплексор
	mux := http.NewServeMux()

	// Обработчик API
	mux.HandleFunc("/api/hello", helloHandler)
	// mux.HandleFunc("/api/data", dataHandler)

	// Обслуживание статических файлов (frontend для VK Mini Apps)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)

	// Настройка CORS для VK Mini Apps
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://vk.com", "https://vk-apps.example.com"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Получаем порт из переменных окружения (для Yandex Cloud / VK Cloud)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Запускаем сервер с CORS
	handler := c.Handler(mux)
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"response": "Hello World!"}`))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Данные из бэкенда!"}`))
}
