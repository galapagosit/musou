package app

import (
	"fmt"
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
	red := color.New(color.FgRed).SprintFunc()
	return red(s)
}
