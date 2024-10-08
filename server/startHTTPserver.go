package main

import (
	"fmt"
	"log"
	"net/http"
)

func startHTTPServer(address string) {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/userdata.json", POSThandler)
	fmt.Printf("HTTP-сервер открыт и прослушивает адрес %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalln("Ошибка запуска HTTP-сервера", err)
	}
}
