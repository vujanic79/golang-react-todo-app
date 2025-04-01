package http_response

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, httpStatus int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling json response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpStatus)
	_, err = w.Write(body)
	if err != nil {
		log.Println("Error writing response", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func RespondWithError(w http.ResponseWriter, httpStatus int, errorMessage string) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, httpStatus, errorResponse{Error: errorMessage})
}
