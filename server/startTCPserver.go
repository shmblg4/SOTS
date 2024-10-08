package main

import (
	"fmt"
	"log"
	"net"
)

func startTCPServer(address string) {
	SERVER, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Ошибка открытия TCP-сервера", err)
	}
	defer SERVER.Close()
	fmt.Printf("TCP-сервер открыт и прослушивает адрес %s\n", address)

	for {
		CONN, err := SERVER.Accept()
		if err != nil {
			log.Println("Ошибка подключения", err)
			continue
		}
		go handleConnection(CONN)
	}
}
