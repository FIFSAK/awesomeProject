package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DB struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}

func healthCheck(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hi i'm Anuar and this app is first go TSIS in this app you can upload file in body and get some info about uploaded file available urls you can see on '/' page")
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "available urls \n'/file' {GET} upload file\n'/file' {GET} return all files\n'/file/{id}' {POST} return exact file")
}

func uploadFile(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)

	file, header, err := request.FormFile("file")
	if err != nil {
		fmt.Fprintf(writer, "Error Retrieving the File\n")
		fmt.Println(err)
		return
	}
	defer file.Close()

	filename := header.Filename
	filesize := header.Size
	filetype := header.Header.Get("Content-Type")

	newData := DB{Name: filename, Type: filetype, Size: filesize}

	fileData, err := ioutil.ReadFile("db.json")

	data := []DB{}

	json.Unmarshal(fileData, &data)

	data = append(data, newData)

	updatedData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	if err := ioutil.WriteFile("db.json", updatedData, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Fprintf(writer, "File Name: %s upload successful", filename)
}

func getFiles(writer http.ResponseWriter, request *http.Request) {
	fileData, _ := ioutil.ReadFile("db.json")

	data := []DB{}

	json.Unmarshal(fileData, &data)
	s := ""
	for i := 0; i < len(data); i++ {
		s += "Name: " + data[i].Name + " Type: " + data[i].Type + " Size: " + strconv.FormatInt(data[i].Size, 10) + " "
	}
	fmt.Println(s)
	fmt.Fprintf(writer, s)
}

func getFileByName(writer http.ResponseWriter, request *http.Request) {

	fileName := mux.Vars(request)["name"]

	fileData, _ := ioutil.ReadFile("db.json")

	data := []DB{}

	json.Unmarshal(fileData, &data)
	s := ""
	for i := 0; i < len(data); i++ {
		if data[i].Name == fileName {
			s += "Name: " + data[i].Name + " Type: " + data[i].Type + " Size: " + strconv.FormatInt(data[i].Size, 10)
		}
	}
	fmt.Println(fileName)
	fmt.Fprintf(writer, s)
}

func deleteFile(writer http.ResponseWriter, request *http.Request) {

	fileName := mux.Vars(request)["name"]

	fileData, _ := ioutil.ReadFile("db.json")

	data := []DB{}

	json.Unmarshal(fileData, &data)
	for i, file := range data {
		if file.Name == fileName {
			data = append(data[:i], data[i+1:]...)
			//break
		}
	}
	updatedData, _ := json.Marshal(data)

	if err := ioutil.WriteFile("db.json", updatedData, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Fprintf(writer, "deleted succes")

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/health-check", healthCheck).Methods("GET")

	router.HandleFunc("/", home)

	router.HandleFunc("/file", uploadFile).Methods("POST")

	router.HandleFunc("/file", getFiles).Methods("GET")

	router.HandleFunc("/file/{name}", getFileByName).Methods("GET")

	router.HandleFunc("/file/{name}", deleteFile).Methods("Delete")
	http.ListenAndServe(":8080", router)
}
