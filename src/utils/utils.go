package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Errors []string    `json:"errors"`
}

func WriteJSON(w http.ResponseWriter, code int, data interface{}, errs []error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var errStrs []string
	for _, err := range errs {
		errStrs = append(errStrs, err.Error())
	}

	resp := APIResponse{
		Code:   code,
		Data:   data,
		Errors: errStrs,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
