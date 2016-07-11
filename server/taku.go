package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"github.com/galapagosit/musou/common"
	"strconv"
)

type Taku struct {
	room_id string
	members []*Member
	yama    []string
}

func NewTaku(room_id string) *Taku {
	taku := new(Taku)
	taku.room_id = room_id
	taku.yama = common.MakeYama()
	ShuffleYama(taku.yama)
	return taku
}

func (taku *Taku)AddMember(member *Member) {
	fmt.Println("member add:", member)
	taku.members = append(taku.members, member)
	member.room_id = taku.room_id
	message := "room_id:" + taku.room_id + " count:" + strconv.Itoa(len(taku.members))
	for _, member := range taku.members {
		websocket.Message.Send(member.ws, message)
	}
}

func (taku *Taku)SaySomething(member *Member, str string) {
	fmt.Println("member say:", member, str)
	for _, member := range taku.members {
		websocket.Message.Send(member.ws, str)
	}
}

