package main

import (
	"nomadproject.io/demo/connect-native/internal/frontend"
	"nomadproject.io/demo/connect-native/internal/options"
)

func main() {
	frontend.Start(
		options.Environment("uuid-fe"),
		options.Consul(),
	).Wait()
}
