package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"volume-assignment/pkg/calculate"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) CalculateResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var pairs [][]string
	err := json.NewDecoder(r.Body).Decode(&pairs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := calculate.FindPath(pairs)
	fmt.Println(pairs)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
