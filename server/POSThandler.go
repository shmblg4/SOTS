package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func POSThandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request struct {
			Action string   `json:"action"`
			User   UserData `json:"user"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
			return
		}

		dataMutex.Lock()
		defer dataMutex.Unlock()

		data, err := loadUserDataFromFile()
		if err != nil {
			http.Error(w, "Ошибка загрузки базы данных", http.StatusInternalServerError)
			return
		}

		switch request.Action {
		case "register":
			if _, exists := data[request.User.Login]; exists {
				http.Error(w, "Пользователь уже существует", http.StatusConflict)
				return
			}
			data[request.User.Login] = request.User
			if err := saveUserDataToFile(data); err != nil {
				http.Error(w, "Ошибка сохранения записи в базу данных", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Printf("Пользователь %s успешно добавлен\n", request.User.Login)

		case "login":
			existingUser, exist := data[request.User.Login]
			if !exist {
				http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
				return
			}
			if existingUser.Password != request.User.Password {
				http.Error(w, "Неверный пароль", http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusOK)

		default:
			http.Error(w, "Некорректное действие", http.StatusBadRequest)
		}
		return
	}
	http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
}
