package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	// socket is websocket for this client
	socket *websocket.Conn
	//send is channel on which channel as send
	send chan []byte
	// room is the room this client is chatting in
	room *room
}
