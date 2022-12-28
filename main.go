package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", Calculate)
	http.ListenAndServe(":8080", nil)
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	// { "a": 1, "b": 2}
	type Request struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	result := Plus(req.A, req.B)
	resp := struct {
		Result int `json:"result"`
	}{
		Result: result,
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
