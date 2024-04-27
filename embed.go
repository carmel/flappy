package main

import (
	_ "embed"
)

var (
	//go:embed yo.ogg
	yo []byte

	//go:embed a.wav
	a []byte

	//go:embed logo.jpeg
	logo []byte
)
