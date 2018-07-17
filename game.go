package bingogame

func NewNumberBox(number int) []int {
	return Shuffle(1, number)
}
