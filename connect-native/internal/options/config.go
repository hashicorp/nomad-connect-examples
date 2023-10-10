// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package options

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/consul/api"
)

type Config struct {
	// Name of this service as registered in Consul.
	Name string

	// Listen is the address & port that this service is going to listen to.
	Listen string

	// Upstream (utilized by frontend) is the Name of the API service in Consul.
	Upstream string
}

func mustGetInt(name string) int {
	if s := os.Getenv(name); s != "" {
		p, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(name + " must be a number")
		}
		return p
	}
	log.Fatal(name + " must be set")
	return -1
}

func getStringOr(name, alt string) string {
	if s := os.Getenv(name); s != "" {
		return s
	}
	return alt
}

func Environment(name string) *Config {
	port := mustGetInt("PORT")
	bind := getStringOr("BIND", "localhost")
	return &Config{
		Name:     name,
		Listen:   fmt.Sprintf("%s:%d", bind, port),
		Upstream: os.Getenv("UPSTREAM"),
	}
}

func Consul() *api.Client {
	logEnvironment("CONSUL_HTTP_ADDR")
	logEnvironment("CONSUL_NAMESPACE")
	logEnvironment("CONSUL_CACERT")
	logEnvironment("CONSUL_CLIENT_CERT")
	logEnvironment("CONSUL_CLIENT_KEY")
	logEnvironment("CONSUL_HTTP_SSL")
	logEnvironment("CONSUL_HTTP_SSL_VERIFY")
	logEnvironment("CONSUL_TLS_SERVER_NAME")
	logEnvironment("CONSUL_GRPC_ADDR")
	logEnvironment("CONSUL_HTTP_TOKEN_FILE")
	consulConfig := api.DefaultConfig()
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal("failed to make consul client:", err)
	}
	return consulClient
}

func logEnvironment(name string) {
	value := os.Getenv(name)
	if value == "" {
		value = "<unset>"
	}
	log.Printf("environment %s = %s", name, value)
}
