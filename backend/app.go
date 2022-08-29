package main

import "fmt"

type Payload map[string]interface{}

func start(ipc IPCInterface) {
	//your app code goes here
	payload := make(map[string]interface{})
	payload["x"] = "vibes"
	ipc.Emit("message", payload)
	ipc.On("message", guy)
}

func guy(p Payload) {
	a := p["val"].(map[string]interface{})

	fmt.Println(a)
}
