// +build tools

// Package tools tracks dependencies for tools that used in the build process.
// See https://github.com/golang/go/issues/25922
// for test ci
package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
