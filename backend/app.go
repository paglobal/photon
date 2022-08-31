package main

import "fmt"

type Payload struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func start() {
	//your app code goes here
	payload := Payload{
		"How you doing?",
		"",
	}

	ipcHub.On("add", func(ipcID string) {
		fmt.Println("New")
		ipc := ipcHub.GetIPC(ipcID)
		payload.ID = ipc.ID
		ipc.Emit("message", payload)
		ipc.On("message", guy)
	})

	ipcHub.On("remove", func(ipcID string) {
		fmt.Println(ipcID)
	})
}

func guy(p Payload, ipc *IPC) {
	payload := Payload{
		"How you doing?",
		ipc.ID,
	}

	ipc.Emit("message", payload)
	fmt.Println(p.Message)
}
