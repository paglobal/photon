package main

import (
	"embed"
)

//go:embed dist/static/*
var content embed.FS

var ipcPort string = ":53174"
var fileServerPort string = ":53172"
