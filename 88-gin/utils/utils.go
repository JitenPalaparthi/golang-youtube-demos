package utils

import (
	"demo/models"
	"encoding/json"
	"os"
)

func SaveToFile(filaName string, user *models.User) error {
	file, err := os.OpenFile(filaName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	bytes = append(bytes, 10)
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
