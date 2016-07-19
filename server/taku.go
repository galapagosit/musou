package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	C "github.com/galapagosit/musou/common"
	"strconv"
	"strings"
	"encoding/json"
)

func recvTakuCommand(taku *Taku) {
	for command := range taku.c {
		command_list := strings.Split(command.Command, " ")
		cmd := command_list[0]
		if (cmd == "say") {
			message := command_list[1]
			taku.SaySomething(command.Member, message)
		}
	}
}

type yama []C.Hai

func (p yama) Len() int {
	return len(p)
}

func (p yama) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type members []*Member

func (p members) Len() int {
	return len(p)
}

func (p members) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

const (
	PHASE_TUMO int = iota
	PHASE_WAIT
)

type Sutehai struct {
	hai         C.Hai
	is_tumogiri bool
}

type turn_member_index int

type Taku struct {
	room_id           string
	members           members
	yama              yama
	tehai_map         map[turn_member_index][]C.Hai
	tsumohai_map      map[turn_member_index]C.Hai
	sutehai_map       map[turn_member_index][]Sutehai
	turn_member_index turn_member_index
	phase             int
	c                 chan *MemberCommand
}

func NewTaku(room_id string) *Taku {
	taku := new(Taku)
	taku.room_id = room_id
	taku.yama = C.MakeYama()
	taku.c = make(chan *MemberCommand)

	go recvTakuCommand(taku)
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
	if (len(taku.members) >= 2) {
		taku.Start()
	}
}

func (taku *Taku)SaySomething(member *Member, str string) {
	fmt.Println("member say:", member, str)
	for _, member := range taku.members {
		websocket.Message.Send(member.ws, str)
	}
}

func (taku *Taku)Start() {
	Shuffle(taku.yama)
	Shuffle(taku.members)
	taku.turn_member_index = 0
	taku.phase = PHASE_WAIT
	taku.tehai_map = make(map[turn_member_index][]C.Hai)
	taku.tsumohai_map = make(map[turn_member_index]C.Hai)
	taku.sutehai_map = make(map[turn_member_index][]Sutehai)

	taku.Haipai()
	taku.Tumo()
	taku.SendStats()
}

func (taku *Taku)Tumo() {
	tumos := taku.PopHai(1)
	taku.tsumohai_map[taku.turn_member_index] = tumos[0]
}

func (taku *Taku)GetMemberIndex(member *Member) turn_member_index {
	for i, m := range taku.members {
		if (m == member) {
			return turn_member_index(i)
		}
	}
	panic("can detect index")
}

func (taku *Taku)PopHai(num int) []C.Hai {
	var tumos []C.Hai
	tumos, taku.yama = taku.yama[:num], taku.yama[num:]
	return tumos
}

func (taku *Taku)Haipai() {
	for _, member := range taku.members {
		tumos := taku.PopHai(13)
		index := taku.GetMemberIndex(member)
		taku.tehai_map[index] = tumos
	}
}

type Stat struct {
	Kaze              string `json:"kaze"`
	Tehai             []C.Hai `json:"tehai"`
	Tumohai           C.Hai `json:"tumohai"`
	Sutehai           map[string][]Sutehai  `json:"sutehai"`
	Turn_member_index turn_member_index `json:"turn_member_index"`
	Phase             int `json:"phase"`
}

func (taku *Taku)SendStat(member *Member) {
	var kaze string

	index := taku.GetMemberIndex(member)

	if (index == 0) {
		kaze = "東"
	} else if (index == 1) {
		kaze = "南"
	} else if (index == 2) {
		kaze = "西"
	} else if (index == 3) {
		kaze = "北"
	}

	stat := new(Stat)

	stat.Kaze = kaze
	stat.Tehai = taku.tehai_map[index]
	stat.Tumohai = taku.tsumohai_map[index]
	sutehai := make(map[string][]Sutehai)
	for k, v := range taku.sutehai_map{
		sutehai[string(k)] = v
	}
	stat.Sutehai = sutehai

	stat.Turn_member_index = taku.turn_member_index
	stat.Phase = taku.phase

	b, err := json.Marshal(stat)
	if err != nil {
		panic(err)
	}
	websocket.Message.Send(member.ws, string(b))
}

func (taku *Taku)SendStats() {
	for _, member := range taku.members {
		taku.SendStat(member)
	}
}
