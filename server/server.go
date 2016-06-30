package server

import (
	"github.com/galapagosit/musou/common"
	"fmt"
)

func StartServer(port int) {
	yama := common.MakeYama()
	ShuffleYama(yama)
	for _, hai := range yama {
		fmt.Print(common.ToColored(hai) + " ")
	}
	fmt.Println(port)
}
