// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
