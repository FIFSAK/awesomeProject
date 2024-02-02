package pkg

import (
	"awesomeProject/internal/db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func HealthCheck(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Hi i'm Anuar and this app is first go TSIS in this app you can upload file in body and get some info about uploaded file available urls you can see on '/' page")
	if err != nil {
		return
	}
}

func Home(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "available urls \n'/file' {POST} upload file\n'/file' {GET} return all files\n'/file/name' {POST} return exact file\n'/file/name' {DELETE} delete file by name")
	if err != nil {
		return
	}
}

func UploadFile(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		return
	}

	file, header, err := request.FormFile("file")
	if err != nil {
		_, err := fmt.Fprintf(writer, "Error Retrieving the File\n")
		if err != nil {
			return
		}
		fmt.Println(err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	filename := header.Filename
	filesize := header.Size
	filetype := header.Header.Get("Content-Type")

	newData := db.DB{Name: filename, Type: filetype, Size: filesize}

	fileData, err := os.ReadFile("internal/db/db.json")

	data := []db.DB{}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return
	}

	data = append(data, newData)

	updatedData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	if err := os.WriteFile("internal/db/db.json", updatedData, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	_, err = fmt.Fprintf(writer, "File Name: %s upload successful", filename)
	if err != nil {
		return
	}
}

func GetFiles(writer http.ResponseWriter, request *http.Request) {
	fileData, _ := os.ReadFile("internal/db/db.json")

	var data []db.DB

	err := json.Unmarshal(fileData, &data)
	if err != nil {
		return
	}
	s := ""
	for i := 0; i < len(data); i++ {
		s += "Name: " + data[i].Name + " Type: " + data[i].Type + " Size: " + strconv.FormatInt(data[i].Size, 10) + " "
	}
	fmt.Println(s)
	_, err = fmt.Fprintf(writer, s)
	if err != nil {
		return
	}
}

func GetFileByName(writer http.ResponseWriter, request *http.Request) {

	fileName := mux.Vars(request)["name"]

	fileData, _ := os.ReadFile("internal/db/db.json")

	data := []db.DB{}

	err := json.Unmarshal(fileData, &data)
	if err != nil {
		return
	}
	s := ""
	for i := 0; i < len(data); i++ {
		if data[i].Name == fileName {
			s += "Name: " + data[i].Name + " Type: " + data[i].Type + " Size: " + strconv.FormatInt(data[i].Size, 10)
		}
	}
	fmt.Println(fileName)
	_, err = fmt.Fprintf(writer, s)
	if err != nil {
		return
	}
}

func DeleteFile(writer http.ResponseWriter, request *http.Request) {

	fileName := mux.Vars(request)["name"]

	fileData, _ := os.ReadFile("internal/db/db.json")

	data := []db.DB{}

	err := json.Unmarshal(fileData, &data)
	if err != nil {
		return
	}
	for i, file := range data {
		if file.Name == fileName {
			data = append(data[:i], data[i+1:]...)
			//break
		}
	}
	updatedData, _ := json.Marshal(data)

	if err := os.WriteFile("internal/db/db.json", updatedData, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	_, err = fmt.Fprintf(writer, "deleted succes")
	if err != nil {
		return
	}
}
