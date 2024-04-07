//go:build windows && amd64

package zyembed

import (
	_ "embed"
)

//go:embed windows-amd64.dll
var lib []byte

func init() { Data = lib }
