package controller

import (
	"AssignmentThree_2nd/.idea/src/entity"
	"bytes"
	"encoding/json"
	"net/http"
)

func orderHandler(w http.ResponseWriter, r *http.Request) {
	order := entity.order{}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create request for the paymentHandler
	requestBody, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call api to service payment to check the amount
	response, err := http.Post("http://10.82.71.188:1521/payment", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// parse response from paymentHandler
	paymentResponse := make(map[string]string)
	err = json.NewDecoder(response.Body).Decode(&paymentResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// response for orderHandler
	responseBody := make(map[string]string)
	responseBody["status"] = paymentResponse["status"]
	if paymentResponse["status"] == "success" {
		responseBody["message"] = "Order successfully"
	} else {
		responseBody["message"] = "Failed to order: " + paymentResponse["message"]
	}

	// encode response body as JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
