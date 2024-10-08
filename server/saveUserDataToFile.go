// saveUserDataToFile.go
package main

import (
	"encoding/json"
	"os"
)

func saveUserDataToFile(data map[string]UserData) error {
	file, err := os.OpenFile("./static/userdata.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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
