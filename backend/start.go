package main

import (
	"fmt"

	"github.com/promethiumjs/photon-lib-go/photon"
)

func start() {
	//your app code goes here
	payload := make(map[string]interface{})
	payload["message"] = "How you doing?"

	photon.IPCHub.On("add", func(ipcID string) {
		fmt.Println("New")
		ipc := photon.IPCHub.GetIPC(ipcID)
		payload["id"] = ipc.ID
		ipc.Emit("message", payload)
		ipc.On("message", printMessage)
	})

	photon.IPCHub.On("remove", func(ipcID string) {
		fmt.Println(ipcID)
	})
}

type MessageInfo struct {
	Message string `json:"message"`
}

func printMessage(p photon.Payload, e string, ipc *photon.IPC) {
	payload := make(map[string]interface{})
	payload["message"] = "How you doing?"
	payload["id"] = ipc.ID

	ipc.Emit("message", payload)
	messageInfo := MessageInfo{}
	photon.ToStruct(p, &messageInfo)
	fmt.Println(p)
	fmt.Println(messageInfo)
}
