package common

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
