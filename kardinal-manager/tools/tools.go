//go:build tools
// +build tools

package main

// It follows the `tools.go` pattern described here: https://github.com/deepmap/oapi-codegen?tab=readme-ov-file#install so we can run the codegen without the need to install the binary
import (
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"
)
