package bingogame

import (
	"math/rand"
	"time"
)

func Shuffle(startNumber, endNumber int) []int {
	length := endNumber - startNumber + 1
	shuffle := make([]int, length)
	for index := 0; index < endNumber; index++ {
		shuffle[index] = index + 1
	}
	for index := len(shuffle) - 1; index > 0; index-- {
		rand.Seed(time.Now().UTC().UnixNano())
		randomIndex := rand.Intn(index + 1)
		shuffle[index], shuffle[randomIndex] = shuffle[randomIndex], shuffle[index]
	}

	return shuffle

}
