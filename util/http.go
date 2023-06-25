package util

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, s int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		_, _ = w.Write([]byte("test"))
	}
}
