package main

import (
	"fmt"
	"github.com/galapagosit/musou/app"
)

func main() {
	yama := app.MakeYama()
	app.ShuffleYama(yama)
	for _, hai := range yama {
		fmt.Print(app.ToColored(hai))
		fmt.Print(" ")
	}
}
