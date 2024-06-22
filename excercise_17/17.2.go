package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	l := log.New(os.Stdout, "", 0)
	logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: logHandler(authHandler(mux)),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		l.Fatalf("Не удалось запустить сервер: %v", err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)
}

// Middleware для авторизации
func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("X-My-App-Id")
		if xId != "my_secret" {
			errorMsg := "пользователь не авторизован"
			http.Error(w, errorMsg, http.StatusUnauthorized)
			log.Println("error: %s", errorMsg)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url: %s", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
