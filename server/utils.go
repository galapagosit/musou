package server

import (
	"crypto/rand"
	"math/big"
)

func GetRandInt(max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return int(n)
}

type Shuffled interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(data Shuffled) {
	n := data.Len()
	for i := 0; i < n; i++ {
		j := GetRandInt(i + 1)
		data.Swap(i, j)
	}
}
