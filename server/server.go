package server

import (
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
	"github.com/galapagosit/musou/common"
	"strings"
)

type Command struct {
	Member *Member `json:"member"`
	Str    string `json:"str"`
}

func recvCommand(recvChan <-chan *Command) {
	var taku Taku;
	for command := range recvChan {
		if (strings.HasPrefix(command.Str, "join ")) {
			taku.AddMember(command.Member)
		} else {
			taku.SaySomething(command.Member, command.Str)
		}
	}
}

func makeEchoHandler() func(ws *websocket.Conn) {
	return func(ws *websocket.Conn) {
		fmt.Println("connected!")

		recvChan := make(chan *Command)
		go recvCommand(recvChan)

		member := &Member{ws: ws}
		for {
			var str string
			if err := websocket.Message.Receive(ws, &str); err != nil {
				fmt.Println("err: ", err)
				break
			}
			fmt.Println("chat receive:", str)
			recvChan <- &Command{Member:member, Str:str}
		}
	}
}

func StartServer(port string) {
	yama := common.MakeYama()
	ShuffleYama(yama)
	for _, hai := range yama {
		fmt.Print(common.ToColored(hai) + " ")
	}
	fmt.Println(port)

	echoHandler := makeEchoHandler()
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(echoHandler)}
			s.ServeHTTP(w, req)
		})
	err := http.ListenAndServe(":" + port, nil);
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
