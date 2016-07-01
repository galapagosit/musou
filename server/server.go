package server

import (
	"fmt"
	"io"
	"net/http"
	"golang.org/x/net/websocket"
	"github.com/galapagosit/musou/common"
)

func makeEchoHandler() func(ws *websocket.Conn) {
	var taku Taku;
	return func(ws *websocket.Conn) {
		fmt.Println("connected!")
		ws.Write([]byte("hello!"))

		member := &Member{ws: ws}
		taku.AddMember(member)

		var err error;
		var written int64;

		buf := make([]byte, 32 * 1024)
		for {
			nr, er := ws.Read(buf)
			if nr > 0 {
				taku.SaySomething(member, string(buf[0:nr]))
				nw, ew := ws.Write(buf[0:nr])
				if nw > 0 {
					written += int64(nw)
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = io.ErrShortWrite
					break
				}
			}
			if er == io.EOF {
				break
			}
			if er != nil {
				err = er
				break
			}
		}
		fmt.Println(err)
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
