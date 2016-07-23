package client

import (
	"golang.org/x/net/websocket"
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	C "github.com/galapagosit/musou/common"
)


func recv(stat_str string) {
	var stat C.Stat
	err := json.Unmarshal([]byte(stat_str), &stat)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(&stat)
}

func reactor(ws *websocket.Conn, recv_c <-chan string, input_c <-chan string) {
	for {
		select {
		case stat_str := <-recv_c:
			recv(stat_str)
		case input := <-input_c:
			if _, err := ws.Write([]byte(input)); err != nil {
				panic("Write: " + err.Error())
			}
		}
	}
}

func receiver(ws *websocket.Conn, recv_c chan <- string) {
	for {
		var command string
		if err := websocket.Message.Receive(ws, &command); err != nil {
			fmt.Println("err: ", err)
			break
		}
		recv_c <- command
	}
}

func scanner(input_c chan <- string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input_c <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func StartClient(host string, port string) {
	ws_url := fmt.Sprintf("ws://%s:%s", host, port)
	http_url := fmt.Sprintf("http://%s:%s", host, port)

	ws, err := websocket.Dial(ws_url, "", http_url);
	if err != nil {
		panic("Dial: " + err.Error())
	}

	recv_c := make(chan string)
	input_c := make(chan string)

	go reactor(ws, recv_c, input_c)

	go receiver(ws, recv_c)
	scanner(input_c)
}
