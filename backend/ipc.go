package main

import (
	"log"
	"reflect"

	"github.com/gorilla/websocket"
)

type Callback func(payload Payload, ipc *IPC)

type EventsMap map[string][]Callback

type IPC struct {
	OnEvents   EventsMap
	OnceEvents EventsMap
	Socket     *websocket.Conn
	ID         string
}

type Data struct {
	Event   string  `json:"event"`
	Payload Payload `json:"payload"`
}

func (ipc *IPC) On(event string, callback Callback) func() {
	return ipc.RegisterEvent(event, callback, "on")
}

func (ipc *IPC) Once(event string, callback Callback) func() {
	return ipc.RegisterEvent(event, callback, "once")
}

func (ipc *IPC) ReturnEventsMap(t string) EventsMap {
	if t == "on" {
		return ipc.OnEvents
	} else {
		return ipc.OnceEvents
	}
}

func (ipc *IPC) RegisterEvent(event string, callback Callback, t string) func() {
	var eventsMap EventsMap
	if t == "on" {
		eventsMap = ipc.OnEvents
	} else {
		eventsMap = ipc.OnceEvents
	}

	if _, ok := eventsMap[event]; !ok {
		var eventArray []Callback
		eventsMap[event] = eventArray
	}

	eventsMap[event] = append(eventsMap[event], callback)

	return func() {
		for i, v := range eventsMap[event] {
			p1 := reflect.ValueOf(v).Pointer()
			p2 := reflect.ValueOf(callback).Pointer()
			if p1 == p2 {
				eventsMap[event] = remove(eventsMap[event], i)
				if len(eventsMap[event]) == 0 {
					delete(eventsMap, event)
				}

				return
			}
		}
	}
}

func (ipc *IPC) Emit(event string, payload Payload) {
	data := Data{event, payload}

	if err := ipc.Socket.WriteJSON(data); err != nil {
		log.Println(err)
	}
}
