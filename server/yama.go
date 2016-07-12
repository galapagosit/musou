package server

func ShuffleYama(yama []string) {
	for i := range yama {
		j := GetRandInt(i + 1)
		yama[i], yama[j] = yama[j], yama[i]
	}
}
