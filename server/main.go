// main.go
package main

func main() {
	addressTCP := "185.204.0.114:8000"
	addressHTTP := "185.204.0.114:8080"
	go startTCPServer(addressTCP)
	go startHTTPServer(addressHTTP)
	select {}
}
