package main

import (
	"fmt"
	"github.com/galapagosit/musou/common"
	"github.com/galapagosit/musou/server"
)

func main() {
	yama := common.MakeYama()
	server.ShuffleYama(yama)
	for _, hai := range yama {
		fmt.Print(common.ToColored(hai))
		fmt.Print(" ")
	}
}
