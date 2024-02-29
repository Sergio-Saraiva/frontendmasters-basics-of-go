package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"frontendmasters.com/courses/go-basics/fem-museum/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var inputExhibition data.Exhibition
	err := json.NewDecoder(r.Body).Decode(&inputExhibition)
	if err != nil {
		fmt.Println("Error deconding body", err)
		http.Error(w, "Error deconding body", http.StatusBadRequest)
		return
	}

	data.Add(inputExhibition)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ok"))
}
