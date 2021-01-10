package frontend

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"nomadproject.io/demo/connect-native/internal/common"
)

const (
	// When using Consul DNS features, this DNS format automatically resolves
	// for the desired connect-enabled consul service.
	//
	// https://www.consul.io/docs/agent/dns#connect-capable-service-lookups
	upstreamURL = "https://%s.service.consul/"
)

type response struct {
	UUID string
}

func handler(upstream string, client *http.Client) http.Handler {
	router := mux.NewRouter()
	router.Handle("/health", common.NewHealthEP())
	router.Handle("/", index(upstream, client))
	return router
}

func index(upstream string, client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uuid, err := getUUID(upstream, client)
		if err != nil {
			log.Println("failed to reach api service:", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		// query the API for a generated UUID
		if err := indexHTML.Execute(w, response{UUID: uuid}); err != nil {
			log.Fatal("failed to render template:", err)
		}
	}
}

func getUUID(upstream string, client *http.Client) (string, error) {
	response, err := client.Get(fmt.Sprintf(
		upstreamURL,
		upstream,
	))
	if err != nil {
		return "", errors.Wrap(err, "failed to reach upstream")
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response from upstream")
	}

	return string(b), nil
}
