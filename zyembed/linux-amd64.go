//go:build amd64 && !windows

package zyembed

import (
	_ "embed"
)

//go:embed linux-amd64.so
var lib []byte

func init() { Data = lib }
