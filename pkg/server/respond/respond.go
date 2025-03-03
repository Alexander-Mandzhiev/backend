package respond

import (
	sl "backend/pkg/logger"
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Key   string            `json:"key"`
	Value map[string]string `json:"value"`
}

func Respond(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "VMK-ds")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		sl.Log.Error("failed to encode response", err)
		return
	}
	return
}

func RespondedError(w http.ResponseWriter, r *http.Request, code int, err error) {
	sl.Log.Error("error occurred", sl.Err(err))
	Respond(w, r, code, &ErrorResponse{
		Key:   "error",
		Value: map[string]string{"message": err.Error()},
	})
	return
}
