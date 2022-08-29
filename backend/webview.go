//go:build webview
// +build webview

package main

import (
	"os"
	"time"

	"github.com/webview/webview"
)

func init() {
	//serve files only when not in dev mode
	args := os.Args[1:]
	if argsCount := len(args); argsCount <= 0 {
		go serve()
	}

	go ipcInit()

	time.Sleep(2 * time.Second)
	launchUI()
}

//launch webview to display frontend UI
func launchUI() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Hello Photon")
	w.SetSize(480, 320, webview.HintNone)
	w.Navigate("http://127.0.0.1:5173/")
	w.Run()
}
