// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generator

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid2 "github.com/hashicorp/go-uuid"
	"nomadproject.io/demo/connect-native/internal/common"
)

func handler() http.Handler {
	router := mux.NewRouter()
	router.Handle("/health", common.NewHealthEP())
	router.Handle("/", generate())
	return router
}

func generate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid, err := uuid2.GenerateUUID()
		if err != nil {
			log.Panic("failed to generate uuid:", err)
		}
		_, _ = io.WriteString(w, uuid)
	}
}
