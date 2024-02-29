package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"frontendmasters.com/courses/go-basics/fem-museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	if id == nil {
		fmt.Println("Error getting the ID")
		http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		return
	}

	finalId, err := strconv.Atoi(id[0])
	if err != nil {
		fmt.Println("Error converting the id from string to int", err)
		http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(data.GetAll()[finalId])
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.GetAll())
}
