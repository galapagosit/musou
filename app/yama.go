package app

import (
	"crypto/rand"
	"math/big"
)

func getRandInt(max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
    if err != nil {
        panic(err)
    }
    n := nBig.Int64()
	return int(n)
}

func ShuffleYama(yama []string) {
	for i := range yama {
		j := getRandInt(i + 1)
		yama[i], yama[j] = yama[j], yama[i]
	}
}
