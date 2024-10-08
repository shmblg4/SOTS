// processData.go
package main

import "net"

func processData(userData UserData, Conn net.Conn) (error, int) {
	userDataMap, err := loadUserDataFromFile()
	if err != nil {
		return err, 0
	}

	dataMutex.Lock()
	defer dataMutex.Unlock()

	existingUser, ok := userDataMap[userData.Login]
	if ok {
		if existingUser.Password != userData.Password {
			Conn.Write([]byte("Неверный пароль"))
			return nil, 1
		}
		for _, newSignal := range userData.Signals {
			existingUser.Signals = append(existingUser.Signals, newSignal)
		}
		userDataMap[userData.Login] = existingUser
	} else {
		// Пользователь не существует, создаем новую запись с новым сигналом
		userDataMap[userData.Login] = userData
	}

	err = saveUserDataToFile(userDataMap)
	if err != nil {
		return err, 0
	}

	return nil, 0
}
