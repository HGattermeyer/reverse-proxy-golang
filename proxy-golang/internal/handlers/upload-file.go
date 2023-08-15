package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"proxy-golang/internal/data"
	"proxy-golang/internal/models"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upload a file")
	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var payload []models.ServerFile
	err = json.Unmarshal(buf.Bytes(), &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	data.ServerSlice = append(data.ServerSlice, payload...)

	fmt.Println(data.ServerSlice)

}
