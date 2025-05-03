package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Struct mainan
type Toy struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Brand string `json:"brand"`
}

// Dummy data
var toys = []Toy{
	{ID: 1, Nama: "Gundam RX-78-2", Brand: "Bandai"},
	{ID: 2, Nama: "Nendoroid Miku", Brand: "Good Smile Company"},
}

// Ambil semua mainan
func getToys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toys)
}

// Tambah mainan baru
func addToy(w http.ResponseWriter, r *http.Request) {
	var newToy Toy
	err := json.NewDecoder(r.Body).Decode(&newToy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newToy.ID = len(toys) + 1
	toys = append(toys, newToy)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newToy)
}

// Detail & hapus mainan by ID
func toyDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/toys/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		for _, toy := range toys {
			if toy.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(toy)
				return
			}
		}
		http.NotFound(w, r)

	case http.MethodDelete:
		for i, toy := range toys {
			if toy.ID == id {
				toys = append(toys[:i], toys[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.NotFound(w, r)

	default:
		http.Error(w, "Metode tidak didukung", http.StatusMethodNotAllowed)
	}
}

// Router utama
func main() {
	http.HandleFunc("/toys", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getToys(w, r)
		case http.MethodPost:
			addToy(w, r)
		default:
			http.Error(w, "Metode tidak didukung", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/toys/", toyDetailHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
