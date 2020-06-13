package httprespond

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func With(w http.ResponseWriter, x interface{}) {
	b, err := json.Marshal(x)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}

func WithError(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := errorResponse{
		Message: message,
	}

	b, err := json.Marshal(errorResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(b)
}
