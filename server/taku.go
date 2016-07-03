package server

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Taku struct {
	members []*Member
}

func (taku *Taku)AddMember(member *Member) {
	taku.members = append(taku.members, member)
}

func (taku *Taku)SaySomething(member *Member, str string) {
	fmt.Println(member, "say", str)
	for _, member := range taku.members {
		websocket.Message.Send(member.ws, str)
	}
}
