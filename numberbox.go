package bingogame

func NewNumberBox(number int) [75]int {
	var numberInBox [75]int
	for i := 0; i < number; i++ {
		numberInBox[i] = i + 1
	}
	return numberInBox
}
