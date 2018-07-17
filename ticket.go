package bingogame

func NewTicket(numberOfGrid int) Ticket {
	grid := make([][]State, numberOfGrid)
	for index := 0; index < numberOfGrid; index++ {
		grid[index] = make([]State, numberOfGrid)
	}
	return Ticket{
		SizeX: numberOfGrid,
		SizeY: numberOfGrid,
		Grid:  grid,
	}
}
