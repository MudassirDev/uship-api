package web

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, err error, message string) {
	log.Printf("error while processing request: %v\n", err)

	respondWithJSON(w, code, struct {
		Msg string `json:"msg"`
	}{
		Msg: message,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to encode message"))
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}
