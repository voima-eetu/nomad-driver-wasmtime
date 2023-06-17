// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/c16a/nomad-driver-wasmtime/wasmtime"
	"github.com/hashicorp/go-hclog"

	"github.com/hashicorp/nomad/plugins"
)

func main() {
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of a nomad driver plugin
func factory(log hclog.Logger) interface{} {
	return wasmtime.NewPlugin(log)
}
