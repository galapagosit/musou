package server

import (
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
	"strings"
)

type MemberCommand struct {
	Member  *Member `json:"member"`
	Command string `json:"command"`
}

func recvCommand(c <-chan *MemberCommand) {
	var taku Taku;
	for command := range c {
		if (strings.HasPrefix(command.Command, "join ")) {
			taku.AddMember(command.Member)
		} else {
			taku.SaySomething(command.Member, command.Command)
		}
	}
}

func makeHandler() func(ws *websocket.Conn) {
	c := make(chan *MemberCommand)
	go recvCommand(c)
	return func(ws *websocket.Conn) {
		fmt.Println("connected:", ws)
		member := &Member{ws: ws}
		for {
			var command string
			if err := websocket.Message.Receive(ws, &command); err != nil {
				fmt.Println("err: ", err)
				break
			}
			fmt.Println("command receive:", member, command)
			c <- &MemberCommand{Member:member, Command:command}
		}
	}
}

func StartServer(port string) {
	handler := makeHandler()
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(handler)}
			s.ServeHTTP(w, req)
		})
	err := http.ListenAndServe(":" + port, nil);
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
