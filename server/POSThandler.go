package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func POSThandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newUser UserData
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		dataMutex.Lock()
		defer dataMutex.Unlock()

		if err := decoder.Decode(&newUser); err != nil {
			http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
			return
		}
		data, err := loadUserDataFromFile()
		if err != nil {
			log.Println("Ошибка загрузки базы данных")
		}

		data[newUser.Login] = newUser
		err = saveUserDataToFile(data)

		if err != nil {
			log.Println("Ошибка сохранения записи в базу данных")
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Пользователь %s успешно добавлен", newUser.Login)
		return
	}

	http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
}
