package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func loadUserData() map[string]UserData {
	log.Println("Загрузка данных пользователей из файла userdata.json")
	file, err := os.Open("userdata.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Файл userdata.json не найден, создается пустой набор данных")
			return make(map[string]UserData)
		}
		log.Println("Ошибка чтения файла userdata.json:", err)
		return make(map[string]UserData)
	}
	defer file.Close()

	data := make(map[string]UserData)
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Println("Ошибка декодирования данных из файла userdata.json:", err)
	}
	return data
}

func saveUserDataToFile(data map[string]UserData) error {
	file, err := os.OpenFile("./userdata.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func postDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен запрос POST на /data")
	var requestData struct {
		Action string   `json:"action"`
		User   UserData `json:"user"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println("Ошибка декодирования данных из запроса POST:", err)
		http.Error(w, "Ошибка декодирования данных", http.StatusBadRequest)
		return
	}

	log.Printf("Действие: %s, Логин пользователя: %s\n", requestData.Action, requestData.User.Login)

	dataMutex.Lock()
	defer dataMutex.Unlock()
	data := loadUserData()

	switch requestData.Action {
	case "login":
		log.Println("Попытка входа пользователя")
		existingUser, exists := data[requestData.User.Login]
		if !exists || existingUser.Password != requestData.User.Password {
			log.Println("Ошибка входа: Неверный логин или пароль")
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		}
		log.Println("Успешный вход пользователя")
		w.WriteHeader(http.StatusOK)

	case "register":
		log.Println("Попытка регистрации нового пользователя")
		if _, exists := data[requestData.User.Login]; exists {
			log.Println("Ошибка регистрации: Пользователь уже существует")
			http.Error(w, "Пользователь уже существует", http.StatusConflict)
			return
		}
		data[requestData.User.Login] = requestData.User
		saveerr := saveUserDataToFile(data)
		if saveerr != nil {
			log.Println("Ошибка сохранения данных в файл userdata.json:", saveerr)
			http.Error(w, "Ошибка сохранения данных", http.StatusInternalServerError)
			return
		}
		log.Println("Регистрация прошла успешно")
		w.WriteHeader(http.StatusCreated)

	default:
		log.Println("Ошибка: Неизвестное действие")
		http.Error(w, "Неизвестное действие", http.StatusBadRequest)
	}
}

func loadDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен запрос GET на /loadData")
	username := r.URL.Query().Get("username")
	if username == "" {
		log.Println("Ошибка: Не указан логин пользователя")
		http.Error(w, "Не указан логин", http.StatusBadRequest)
		return
	}

	log.Printf("Загрузка данных для пользователя: %s\n", username)
	data := loadUserData()
	user, exists := data[username]
	if !exists {
		log.Printf("Ошибка: Пользователь %s не найден\n", username)
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	log.Printf("Данные пользователя %s успешно загружены\n", username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Signals)
}

func tcpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен запрос POST на /tcp")
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var userData UserData
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		log.Println("Ошибка декодирования данных в запросе TCP:", err)
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	log.Printf("Получены данные от пользователя %s\n", userData.Login)
	existingUser, exists := loadUserData()[userData.Login]
	if !exists {
		log.Printf("Создание новой записи для пользователя %s\n", userData.Login)
		existingUser.Signals = append(existingUser.Signals, userData.Signals...)
		data := loadUserData()
		data[userData.Login] = userData
		saveerr := saveUserDataToFile(data)
		if saveerr != nil {
			log.Println("Ошибка сохранения данных в файл userdata.json:", saveerr)
			http.Error(w, "Ошибка сохранения данных", http.StatusInternalServerError)
			return
		}
		log.Printf("Запись успешно сохранена\n")
	} else if exists && existingUser.Password != userData.Password {
		log.Println("Ошибка: Неверный пароль")
		http.Error(w, "Неверный пароль", http.StatusUnauthorized)
		return
	} else {
		log.Printf("Обновление данных пользователя %s\n", userData.Login)
		existingUser.Signals = append(existingUser.Signals, userData.Signals...)
		data := loadUserData()
		data[userData.Login] = existingUser
		saveerr := saveUserDataToFile(data)
		if saveerr != nil {
			log.Println("Ошибка сохранения данных в файл userdata.json:", saveerr)
			http.Error(w, "Ошибка сохранения данных", http.StatusInternalServerError)
			return
		}
		log.Printf("Данные пользователя %s успешно обновлены\n", userData.Login)
	}

	log.Printf("Данные пользователя %s успешно обновлены\n", userData.Login)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/data", postDataHandler)
	http.HandleFunc("/loadData", loadDataHandler)
	http.HandleFunc("/tcp", tcpHandler)

	log.Println("Database-service запущен на 0.0.0.0:8081")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
