package handlers

import (
	"demo/models"
	"encoding/json"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

type Userhandler struct{}

func (uh *Userhandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("something went wrong with the request data"))
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := new(models.User)
	err = json.Unmarshal(bytes, user)
	if err != nil {
		w.Write([]byte("something went wrong with the request data"))
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Store the user in the file

	user.Id = uint(rand.IntN(10000)) // just for the demo purpose... usually this should be the serial number
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	err = SaveToFile("users.txt", user)
	if err != nil {
		w.Write([]byte("something went wrong with the request data"))
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("file successfully saved in the system"))
	w.WriteHeader(201)
}

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
