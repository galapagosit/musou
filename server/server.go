package server

import (
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
	"strings"
)

type MemberCommand struct {
	Member  *Member
	Command string
}

var takuMap = make(map[string]*Taku)

func recvCommand(c <-chan *MemberCommand) {
	for command := range c {
		command_list := strings.Split(command.Command, " ")
		cmd := command_list[0]
		if (cmd == "create") {
			room_id := command_list[1]
			taku := NewTaku(room_id);
			taku.AddMember(command.Member)
			takuMap[room_id] = taku
		} else if (cmd == "join") {
			room_id := command_list[1]
			taku := takuMap[room_id]
			taku.AddMember(command.Member)
		} else {
			taku := takuMap[command.Member.room_id]
			taku.c <- command
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
