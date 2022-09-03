package main

import (
	"fmt"

	"github.com/promethiumjs/photon-lib-go/photon"
)

func start() {
	//your app code goes here
	payload := make(photon.Payload)
	payload["message"] = "How you doing?"

	photon.IPCHubInstance.On("add", func(ipcID string) {
		fmt.Println("New")
		ipc := photon.IPCHubInstance.GetIPC(ipcID)
		payload["id"] = ipc.ID
		ipc.Emit("message", payload)
		ipc.On("message", printMessage)
	})

	photon.IPCHubInstance.On("remove", func(ipcID string) {
		fmt.Println(ipcID)
	})
}

func printMessage(p photon.Payload, ipc *photon.IPC) {
	payload := make(photon.Payload)
	payload["message"] = "How you doing?"
	payload["id"] = ipc.ID

	ipc.Emit("message", payload)
	fmt.Println(p["message"])
}
