package handlers

import (
	"encoding/json"
	"net/http"
)

type healthResp struct {
	Status string `json:"status"`
}

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(healthResp{Status: "ok"})
	}
}
