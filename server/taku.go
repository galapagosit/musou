package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	C "github.com/galapagosit/musou/common"
	"strings"
	"encoding/json"
	"strconv"
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

type Taku struct {
	room_id           string
	members           members
	yama              yama
	tehai_map         map[int][]C.Hai
	sutehai_map       map[int][]C.Sutehai
	tsumohai_map      map[int]C.Hai
	dahai_map         map[int]C.Hai
	turn_member_index int
	can_tumo          bool
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
	//message := "room_id:" + taku.room_id + " count:" + strconv.Itoa(len(taku.members))
	//for _, member := range taku.members {
	//	websocket.Message.Send(member.ws, message)
	//}

	// メンツが揃った
	if (len(taku.members) >= 2) {
		taku.Start()
	}
}

func (taku *Taku)SuteHai(member *Member, hai_index int) {
	fmt.Println("sute hai add:", member, hai_index)
	index := taku.GetMemberIndex(member)

	taku.tehai_map[index] = append(taku.tehai_map[index], taku.tsumohai_map[index])
	hai := taku.tehai_map[index][hai_index]
	sutehai := new(C.Sutehai)
	sutehai.Hai = hai
	if (hai_index == 14){
		sutehai.Is_tumogiri = true
	}else{
		sutehai.Is_tumogiri = false
	}
	taku.sutehai_map[index] = append(taku.sutehai_map[index], *sutehai)
	taku.tehai_map[index] = append(taku.tehai_map[index][:hai_index], taku.tehai_map[index][hai_index + 1:]...)
	tsumohai := new(C.Hai)
	taku.tsumohai_map[index] = *tsumohai
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
	taku.can_tumo = true
	taku.tehai_map = make(map[int][]C.Hai)
	taku.sutehai_map = make(map[int][]C.Sutehai)
	taku.tsumohai_map = make(map[int]C.Hai)
	taku.dahai_map = make(map[int]C.Hai)
	for i, _ := range taku.members {
		taku.tehai_map[i] = make([]C.Hai, 0)
		taku.sutehai_map[i] = make([]C.Sutehai, 0)
		taku.tsumohai_map[i] = ""
		taku.dahai_map[i] = ""
	}

	taku.Haipai()
	taku.Tumo()
	taku.SendStats()
}

func (taku *Taku)Tumo() {
	tumos := taku.PopHai(1)
	taku.tsumohai_map[taku.turn_member_index] = tumos[0]
	taku.can_tumo = false
}

func (taku *Taku)GetMemberIndex(member *Member) int {
	for i, m := range taku.members {
		if (m == member) {
			return i
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

	stat := new(C.Stat)

	stat.Kaze = kaze

	stat.Tehai_map = make(map[string][]C.Hai)
	for k, v := range taku.tehai_map {
		if (k == index){
			stat.Tehai_map[strconv.Itoa(k)] = v
		}else{
			// mask
			tehai := []C.Hai{}
			for i := 0; i < len(v); i++ {
				tehai = append(tehai, C.MASKED_HAI)
			}
			stat.Tehai_map[strconv.Itoa(k)] = tehai
		}
	}

	stat.Sutehai_map = make(map[string][]C.Sutehai)
	for k, v := range taku.sutehai_map{
		stat.Sutehai_map[strconv.Itoa(k)] = v
	}

	stat.Tsumohai_map = make(map[string]C.Hai)
	for k, v := range taku.tsumohai_map{
		if (v != ""){
			if (k == index){
				stat.Tsumohai_map[strconv.Itoa(k)] = v
			}else{
				stat.Tsumohai_map[strconv.Itoa(k)] = C.MASKED_HAI
			}
		}else{
			stat.Tsumohai_map[strconv.Itoa(k)] = C.NO_HAI
		}
	}

	stat.Dahai_map =make(map[string]C.Hai)
	for k, v := range taku.dahai_map{
		stat.Dahai_map[strconv.Itoa(k)] = v
	}

	stat.Turn_member_index = taku.turn_member_index
	stat.Can_tumo = taku.can_tumo

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
