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

	// メンツが揃った
	if (len(taku.members) >= 3) {
		taku.Oyagime()
		taku.Haipai()
	}
}

func (taku *Taku)SaySomething(member *Member, str string) {
	fmt.Println("member say:", member, str)
	for _, member := range taku.members {
		websocket.Message.Send(member.ws, str)
	}
}

func (taku *Taku)Oyagime() {
	for i := range taku.members {
		j := GetRandInt(i + 1)
		taku.members[i], taku.members[j] = taku.members[j], taku.members[i]
	}
}

func (taku *Taku)Tumo(num int) []string {
	var tumos []string
	tumos, taku.yama = taku.yama[:num], taku.yama[num:]
	fmt.Println("remain yama:", strconv.Itoa(len(taku.yama)))
	return tumos
}

func (taku *Taku)Haipai() {
	for i, member := range taku.members {
		tumos := taku.Tumo(13)
		var kaze string

		if (i == 0) {
			kaze = "東"
		} else if (i == 1) {
			kaze = "南"
		} else if (i == 2) {
			kaze = "西"
		} else if (i == 3) {
			kaze = "北"
		}
		websocket.Message.Send(member.ws, "あなたの風は" + kaze)
		websocket.Message.Send(member.ws, strings.Join(tumos, " "))
	}
}
