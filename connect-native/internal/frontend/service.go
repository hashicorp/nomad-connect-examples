// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontend

import (
	"log"
	"net/http"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
	"nomadproject.io/demo/connect-native/internal/options"
)

type Service struct {
}

func Start(cfg *options.Config, consul *api.Client) *Service {
	s := new(Service)

	if cfg.Upstream == "" {
		log.Fatal("UPSTREAM must be set")
	}

	log.Println("starting frontend, listening on:", cfg.Listen)

	// create a connect service for this client, for making a connection
	// to the api upstream
	service, err := connect.NewService(cfg.Name, consul)
	if err != nil {
		log.Fatal("failed to establish connect service:", err)
	}

	go func() {
		if err := (&http.Server{
			Addr:    cfg.Listen,
			Handler: handler(cfg.Upstream, service.HTTPClient()),
		}).ListenAndServe(); err != nil {
			log.Fatal("failed to listen and serve:", err)
		}
	}()
	return s
}

func (s *Service) Wait() {
	select {
	// serve forever
	}
}
