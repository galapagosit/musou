package common

import (
	"fmt"
	"strings"
	"github.com/fatih/color"
)

type Hai string

const (
	PHASE_TUMO int = iota
	PHASE_WAIT
)

func HaisToStrings(hais []Hai) []string {
	var strings []string;
	for _, hai := range(hais){
		strings = append(strings, ToColored(hai))
	}
	return strings
}

func MakeYama() []Hai {
	yama := []Hai{}

	for j := 0; j < 4; j++ {
		for i := 1; i <= 9; i++ {
			yama = append(yama, Hai(fmt.Sprintf("m%d", i)))
		}

		for i := 1; i <= 9; i++ {
			yama = append(yama, Hai(fmt.Sprintf("p%d", i)))
		}

		for i := 1; i <= 9; i++ {
			yama = append(yama, Hai(fmt.Sprintf("s%d", i)))
		}

		yama = append(yama, "東", "西", "南", "北")

		yama = append(yama, "白", "発", "中")
	}

	return yama
}

func ToColored(s Hai) string {
	if (strings.HasPrefix(string(s), "m")) {
		red := color.New(color.FgRed).SprintFunc()
		return red(s)
	} else if (strings.HasPrefix(string(s), "p")) {
		white := color.New(color.FgWhite).SprintFunc()
		return white(s)
	} else if (strings.HasPrefix(string(s), "s")) {
		green := color.New(color.FgGreen).SprintFunc()
		return green(s)
	} else if (s == "発") {
		green := color.New(color.FgGreen).SprintFunc()
		return green(s)
	} else if (s == "中") {
		red := color.New(color.FgRed).SprintFunc()
		return red(s)
	} else {
		white := color.New(color.FgWhite).SprintFunc()
		return white(s)
	}
}
