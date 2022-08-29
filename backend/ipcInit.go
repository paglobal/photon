package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Host == "127.0.0.1:5317"
	},
}

func ipcInit() {
	router := gin.Default()

	router.GET(
		"/ipc", func(c *gin.Context) {
			socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				log.Println(err)
				return
			}

			onEvents := make(map[string][]Callback)
			onceEvents := make(map[string][]Callback)

			ipcTemp := IPC{
				onEvents,
				onceEvents,
				socket,
			}

			var ipc IPCInterface = &ipcTemp
			start(ipc)

			defer socket.Close()

			for {
				//Read message from browser
				var data Data
				err := socket.ReadJSON(&data)
				if err != nil {
					log.Println(err)
					return
				}

				onEvents := ipc.ReturnEventsMap("on")
				onceEvents := ipc.ReturnEventsMap("once")

				for _, v := range onEvents[data.Event] {
					v(data.Payload)
				}

				for _, v := range onceEvents[data.Event] {
					v(data.Payload)
				}

				delete(onceEvents, data.Event)
			}
		})

	router.Run("127.0.0.1:5317")
}
