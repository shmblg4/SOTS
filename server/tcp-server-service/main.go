package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

func sendToDatabase(data []byte) error {
	url := "http://db-service:8081/tcp"
	log.Println("Отправка данных в database-service:", string(data))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("Ошибка соединения с database-service:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Ошибка сохранения данных в database-service: статус %d\n", resp.StatusCode)
		return fmt.Errorf("ошибка сохранения данных: статус %d", resp.StatusCode)
	}

	log.Println("Данные успешно отправлены в database-service")
	return nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Новое подключение от %s\n", conn.RemoteAddr().String())

	var userData UserData
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(&userData); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
		conn.Write([]byte("Неверный формат данных\n"))
		return
	}

	data, _ := json.Marshal(userData)
	if err := sendToDatabase(data); err != nil {
		conn.Write([]byte("Ошибка при отправке данных\n"))
		return
	}

	conn.Write([]byte("Данные успешно отправлены\n"))
	log.Printf("Данные от пользователя %s успешно обработаны\n", userData.Login)
}

func main() {
	address := "0.0.0.0:8000"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Ошибка запуска TCP-сервера:", err)
	}
	defer listener.Close()
	log.Printf("TCP-сервер запущен на %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка подключения:", err)
			continue
		}
		go handleConnection(conn)
	}
}
