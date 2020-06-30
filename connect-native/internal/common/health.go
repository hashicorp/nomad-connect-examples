package common

import (
	"log"
	"net/http"
)

type healthEP struct{}

func NewHealthEP() http.Handler {
	return new(healthEP)
}

func (h *healthEP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("health requested from:", r.RemoteAddr)
	http.Error(w, "ok", http.StatusOK)
}
