package client

import (
	"golang.org/x/net/websocket"
	"fmt"
	"bufio"
	"os"
)

func reader(ws *websocket.Conn){
	for {
		var command string
		if err := websocket.Message.Receive(ws, &command); err != nil {
			fmt.Println("err: ", err)
			break
		}
		fmt.Println("command receive:", command)
	}
}

func StartClient(host string, port string) {
	ws_url := fmt.Sprintf("ws://%s:%s",host, port)
	http_url := fmt.Sprintf("http://%s:%s",host, port)

	ws, err := websocket.Dial(ws_url, "", http_url);
	if err != nil {
		panic("Dial: " + err.Error())
	}

	go reader(ws)

	scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
		if _, err := ws.Write([]byte(scanner.Text())); err != nil {
			panic("Write: " + err.Error())
		}
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
