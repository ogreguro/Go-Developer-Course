package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello2)

	//проверка и создание файла лога
	file, err := os.OpenFile("excercise_17/http-service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	//log перенаправлен в файл
	log.SetOutput(file)
	
	l := log.New(os.Stdout, "", log.LstdFlags)
	//кастомный логгер перенаправлен в тот же файл
	l.SetOutput(file)

	logHandler := logMiddleware2(l)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: logHandler(authHandler2(mux)),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		l.Fatalf("Не удалось запустить сервер: %v", err)
	}
}

func hello2(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)
}

// Middleware для авторизации
func authHandler2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("X-My-App-Id")
		if xId != "my_secret" {
			errorMsg := "пользователь не авторизован"
			http.Error(w, errorMsg, http.StatusUnauthorized)
			log.Println("error:", errorMsg)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware2(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
