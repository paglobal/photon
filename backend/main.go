package main

import (
	"os"
)

func main() {
	//serve files only when not in dev mode
	args := os.Args[1:]
	if argsCount := len(args); argsCount <= 0 {
		go serve()
	}

	ipcInit()
}
