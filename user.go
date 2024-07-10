package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name     string `json:"name" xml:"name" yaml:"name"`
	Email    string `json:"email" xml:"email" yaml:"email"`
	Message  string `json:"message" xml:"message" yaml:"message"`
	Interest string `json:"interest" xml:"interest" yaml:"interest"`
}

type SuccessfulMessage struct {
	Message string `json:"msg" xml:"msg" yaml:"msg"`
}

func createUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "This route only accepts POST", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	successfulMessage := SuccessfulMessage{Message: fmt.Sprintf("Thank %s, for getting in touch and sharing your interests. We look forward to hearing from you soon.", user.Name)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&successfulMessage)
}
