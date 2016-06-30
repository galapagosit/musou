package common

import (
	"fmt"
	"strings"
	"github.com/fatih/color"
)

func MakeYama() []string {
	yama := []string{}

	for j := 0; j < 4; j++ {
		for i := 1; i <= 9; i++ {
			yama = append(yama, fmt.Sprintf("m%d", i))
		}

		for i := 1; i <= 9; i++ {
			yama = append(yama, fmt.Sprintf("p%d", i))
		}

		for i := 1; i <= 9; i++ {
			yama = append(yama, fmt.Sprintf("s%d", i))
		}

		yama = append(yama, "東", "西", "南", "北")

		yama = append(yama, "白", "発", "中")
	}

	return yama
}

func ToColored(s string) string {
	if (strings.HasPrefix(s, "m")) {
		red := color.New(color.FgRed).SprintFunc()
		return red(s)
	} else if (strings.HasPrefix(s, "p")) {
		white := color.New(color.FgWhite).SprintFunc()
		return white(s)
	} else if (strings.HasPrefix(s, "s")) {
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
