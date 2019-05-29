package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	//dbName := os.Getenv("DB_NAME")
	//dbPass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")

	router:= mux.NewRouter()
	router.HandleFunc("/users",GetUsers).Methods("GET")

	fmt.Println("magic is happening on port 8081")

	log.Fatal(http.ListenAndServe(":8081",router))
	
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	respondWithJSON(writer, http.StatusOK, "OK")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
