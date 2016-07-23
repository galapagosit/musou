package common

import (
	"fmt"
)

type Sutehai struct {
	Hai         Hai
	Is_tumogiri bool
}

type Stat struct {
	Kaze              string `json:"kaze"`
	Tehai             []Hai `json:"tehai"`
	Tumohai           Hai `json:"tumohai"`
	Sutehai           map[string][]Sutehai  `json:"sutehai"`
	Turn_member_index int `json:"turn_member_index"`
	Phase             int `json:"phase"`
}

func (stat *Stat)String() string {
	var sutehai_s string
	for k, v := range stat.Sutehai {
		sutehai_s += fmt.Sprintf("%s %s\n", k, v)
	}

	return fmt.Sprintf(`風:%s
手牌:%s %s
捨て牌:
%s
ターン: %d %d`,
		stat.Kaze,
		HaisToStrings(stat.Tehai), stat.Tumohai,
		sutehai_s,
		stat.Turn_member_index, stat.Phase,
	)
}
