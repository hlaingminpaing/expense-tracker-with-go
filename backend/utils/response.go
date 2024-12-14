// package utils

// import (
// 	"encoding/json"
// 	"net/http"
// )

// func RespondWithError(w http.ResponseWriter, code int, message string) {
// 	RespondWithJSON(w, code, map[string]string{"error": message})
// }

// func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }


package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, map[string]string{"error": message})
}
