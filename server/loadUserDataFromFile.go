// loadUserDataFromFile.go
package main

import (
	"encoding/json"
	"os"
)

func loadUserDataFromFile() (map[string]UserData, error) {
	dataMap := make(map[string]UserData)
	file, err := os.Open("./static/userdata.json")
	if err != nil {
		if os.IsNotExist(err) {
			return dataMap, nil // Файл не существует, возвращаем пустую map
		}
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&dataMap)
	if err != nil {
		return nil, err
	}
	return dataMap, nil
}
