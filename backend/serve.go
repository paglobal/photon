package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed dist/static/*
var content embed.FS

//serve embedded static frontend files
func serve() {
	router := gin.Default()

	static, err := fs.Sub(content, "dist/static")
	if err != nil {
		fmt.Println(err)
	}
	router.StaticFS("/", http.FS(static))

	router.Run("127.0.0.1:5173")
}
