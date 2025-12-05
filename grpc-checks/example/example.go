// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package example

import (
	"context"
	"log"

	ghc "google.golang.org/grpc/health/grpc_health_v1"
)

// Server is a trivial gRPC server that implements the standard grpc.health.v1
// interface.
type Server struct {
}

func New() *Server {
	return new(Server)
}

func (s *Server) Check(ctx context.Context, hcr *ghc.HealthCheckRequest) (*ghc.HealthCheckResponse, error) {
	log.Printf("Check:%s (%s)", hcr.Service, hcr.String())
	return &ghc.HealthCheckResponse{
		Status: ghc.HealthCheckResponse_SERVING,
	}, nil
}

func (s *Server) List(ctx context.Context, hlr *ghc.HealthListRequest) (*ghc.HealthListResponse, error) {
	log.Printf("List (%s)", hlr.String())
	return &ghc.HealthListResponse{
		Statuses: map[string]*ghc.HealthCheckResponse{
			"foo": {Status: ghc.HealthCheckResponse_SERVING},
		},
	}, nil
}

func (s *Server) Watch(hcr *ghc.HealthCheckRequest, hws ghc.Health_WatchServer) error {
	log.Printf("Watch:%s (%s)", hcr.Service, hcr.String())
	return hws.Send(&ghc.HealthCheckResponse{
		Status: ghc.HealthCheckResponse_SERVING,
	})
}
