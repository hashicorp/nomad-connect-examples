package main

import (
	"nomadproject.io/demo/connect-native/internal/generator"
	"nomadproject.io/demo/connect-native/internal/options"
)

func main() {
	generator.Start(
		options.Environment("uuid-api"),
		options.Consul(),
	).Wait()
}
