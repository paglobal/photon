package main

import (
	"embed"

	"github.com/promethiumjs/photon-lib-go/photon"
)

//go:embed dist/static/*
var content embed.FS

func main() {
	photon.Initialize(content, "dist/static", start)

	photon.IPCInit()
}
