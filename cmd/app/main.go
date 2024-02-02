package main

import (
	"awesomeProject/pkg"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/health-check", pkg.HealthCheck).Methods("GET")

	router.HandleFunc("/", pkg.Home).Methods("GET")

	router.HandleFunc("/file", pkg.UploadFile).Methods("POST")

	router.HandleFunc("/file", pkg.GetFiles).Methods("GET")

	router.HandleFunc("/file/{name}", pkg.GetFileByName).Methods("GET")

	router.HandleFunc("/file/{name}", pkg.DeleteFile).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
