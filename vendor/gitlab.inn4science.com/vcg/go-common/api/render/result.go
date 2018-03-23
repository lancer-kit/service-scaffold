package render

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes some response as WriteJSON to the `http.ResponseWriter`.
func WriteJSON(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	js, err := json.MarshalIndent(resp, "", "  ")

	if err != nil {
		http.Error(w, "error while render response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(js)
}

// ServerError renders default http.StatusInternalServerError.
func ServerError(w http.ResponseWriter) {
	WriteJSON(w, http.StatusInternalServerError, ResultServerError)
}
