package main

import (
	"encoding/json"
	"net/http"
)

func GEThandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "Логин не предоставлен", http.StatusBadRequest)
			return
		}

		data, err := loadUserDataFromFile()
		if err != nil {
			http.Error(w, "Ошибка загрузки базы данных", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data[username].Signals)
		return
	}

	http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	return
}
