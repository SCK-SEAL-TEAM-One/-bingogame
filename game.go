package bingogame

func NewNumberBox(number int) []int {
	return Shuffle(1, number)
}

func (g *Game) PickUpNumber() int {
	var pickUpNumber int
	pickUpNumber, g.NumberBox = g.NumberBox[0], g.NumberBox[1:]
	return pickUpNumber
}
