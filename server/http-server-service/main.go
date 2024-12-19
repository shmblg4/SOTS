package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Новое подключение: метод=%s, путь=%s, IP-адрес=%s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func main() {
	address := "0.0.0.0:8080"

	// Хостинг статики
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", logMiddleware(fs))
	log.Println("HTTP-сервер хостит статические файлы из директории './static'")

	// Обработка POST-запроса для регистрации и авторизации
	http.Handle("/request", logMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		var inputData map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
			log.Println("Ошибка декодирования JSON:", err)
			http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
			return
		}

		action := inputData["action"].(string)
		user := inputData["user"].(map[string]interface{})
		data, err := json.Marshal(map[string]interface{}{
			"action": action,
			"user":   user,
		})
		if err != nil {
			log.Println("Ошибка формирования JSON:", err)
			http.Error(w, "Ошибка формирования JSON", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post("http://db-service:8081/data", "application/json", bytes.NewReader(data))
		if err != nil {
			log.Println("Ошибка при обращении к db-service:", err)
			http.Error(w, "Ошибка обращения к базе данных", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})))

	// Обработка GET-запроса для получения сигналов
	http.Handle("/loadData", logMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "Не указан логин", http.StatusBadRequest)
			return
		}

		resp, err := http.Get(fmt.Sprintf("http://db-service:8081/loadData?username=%s", username))
		if err != nil {
			log.Println("Ошибка при получении данных из db-service:", err)
			http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})))

	log.Printf("HTTP-сервер запущен на %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Ошибка запуска HTTP-сервера:", err)
	}
}
