// handleConnection.go
package main

import (
	"encoding/json"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	decoder := json.NewDecoder(conn)
	var userData UserData
	err := decoder.Decode(&userData)
	if err != nil {
		log.Println("Ошибка при декодировании данных пользователя:", err)
		conn.Write([]byte("Неверный формат данных\n"))
		return
	}

	err_flag := 0

	if processData_err, err_type := processData(userData, conn); processData_err != nil {
		err_flag = err_type
		log.Printf("Ошибка при обработке данных от клиента: %v", processData_err)
		_, _ = conn.Write([]byte("Ошибка при обработке данных\n"))
		return
	}

	if err_flag == 0 {
		_, err = conn.Write([]byte("Данные получены и сохранены\n"))
		if err != nil {
			log.Printf("Ошибка при отправке данных клиенту: %v", err)
		}
	}

}
