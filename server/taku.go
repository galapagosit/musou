package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"github.com/galapagosit/musou/common"
	"strconv"
	"strings"
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
	if(len(taku.members) >= 3){
		taku.Haipai()
	}
}

func (taku *Taku)SaySomething(member *Member, str string) {
	fmt.Println("member say:", member, str)
	for _, member := range taku.members {
		websocket.Message.Send(member.ws, str)
	}
}

func (taku *Taku)Tumo(num int) []string{
	var tumos []string
	tumos, taku.yama = taku.yama[:num], taku.yama[num:]
	fmt.Println("remain yama:", strconv.Itoa(len(taku.yama)))
	return tumos
}

func (taku *Taku)Haipai() {
	for _, member := range taku.members {
		tumos := taku.Tumo(13)
		websocket.Message.Send(member.ws, strings.Join(tumos, " "))
	}
}
