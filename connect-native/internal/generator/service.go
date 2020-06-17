package generator

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

	log.Println("starting generator API, listening on", cfg.Listen)

	cs, err := connect.NewService(cfg.Name, consul)
	if err != nil {
		log.Fatal("failed to create connect service:", err)
	}

	go func() {
		if err := (&http.Server{
			Addr:      cfg.Listen,
			TLSConfig: cs.ServerTLSConfig(),
			Handler:   handler(),
		}).ListenAndServeTLS("", ""); err != nil {
			log.Fatal("failed to listen and serve TLS:", err)
		}
	}()

	return s
}

func (s *Service) Wait() {
	select {
	// serve API forever
	}
}
