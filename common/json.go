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
	Tehai_map         map[string][]Hai `json:"tehai_map"`
	Sutehai_map       map[string][]Sutehai `json:"sutehai_map"`
	Tsumohai_map      map[string]Hai `json:"tumohai_map"`
	Dahai_map         map[string]Hai `json:"dahai_map"`
	Turn_member_index int `json:"turn_member_index"`
	Can_tumo          bool `json:"can_tumo"`
}

func (stat *Stat)String() string {
	var tehai_s string
	for k, v := range stat.Tehai_map {
		tehai_s += fmt.Sprintf("%s %s\n", k, HaisToStrings(v))
	}

	var sutehai_s string
	for k, v := range stat.Sutehai_map {
		var sutehai string
		for _, s := range (v) {
			sutehai += fmt.Sprintf("%s ", s)
		}
		sutehai_s += fmt.Sprintf("%s %s\n", k, sutehai)
	}

	var tsumohai_s string
	for k, v := range stat.Tsumohai_map {
		tsumohai_s += fmt.Sprintf("%s %s\n", k, ToColored(v))
	}

	var dahai_s string
	for k, v := range stat.Dahai_map {
		dahai_s += fmt.Sprintf("%s %s\n", k, ToColored(v))
	}

	return fmt.Sprintf(`風:%s

手牌:
%s
捨て牌:
%s
ツモ牌:
%s
打牌:
%s
ターン: %d
ツモ可能: %t`,
		stat.Kaze,
		tehai_s,
		sutehai_s,
		tsumohai_s,
		dahai_s,
		stat.Turn_member_index,
		stat.Can_tumo,
	)
}
