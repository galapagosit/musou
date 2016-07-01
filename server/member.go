package server

import (
	"golang.org/x/net/websocket"
)

type Member struct {
	ws *websocket.Conn
}
