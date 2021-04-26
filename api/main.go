package main

import (
	"log"
	"net/http"
)

func apiResponse(resp http.ResponseWriter, request *http.Request) {
	//response.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	//response.Write([]byte(`{"message":"hello world"}`))
	switch request.Method {
	case "GET":
		response(`{"message": "GET method requested"}`, resp, http.StatusOK)
	case "POST":
		response(`{"message": "POST method requested"}`, resp, http.StatusCreated)
	default:
		response(`{"message": "Cant find method requested"}`, resp, http.StatusNotFound)
	}
}

func response(text string, response http.ResponseWriter, statusCode int) {
	response.WriteHeader(statusCode)
	response.Write([]byte(text))

}

func main() {
	http.HandleFunc("/", apiResponse)
	log.Fatal((http.ListenAndServe(":8080", nil)))
}
