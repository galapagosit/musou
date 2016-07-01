package server

import "fmt"

type Taku struct {
	members []*Member
}

func (taku *Taku)AddMember(member *Member) {
	taku.members = append(taku.members, member)
}

func (taku *Taku)SaySomething(member *Member, str string) {
	fmt.Println(member, "say", str)
	for _, member := range taku.members {
		member.ws.Write([]byte(str))
	}
}
