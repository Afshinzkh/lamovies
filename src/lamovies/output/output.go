package output

import (
	"encoding/json"
	"net/http"
)

// Error ...
func Error(w http.ResponseWriter, code int, msg string) {
	JSON(w, code, map[string]string{"Error ": msg})
}

// JSON ...
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
