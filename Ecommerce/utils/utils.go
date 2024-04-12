package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJOSN(r *http.Request, payLoad any) error {
	if r.Body == nil {
		return fmt.Errorf("invalid request body")
	}

	return json.NewDecoder(r.Body).Decode(payLoad)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
