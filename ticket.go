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

func GenerateTicketNumber(ticket Ticket) Ticket {
	for indexRow := range ticket.Grid {
		startNumber := 1 + (indexRow * 15)
		endNumber := 15 + (indexRow * 15)
		suffleNumber := Shuffle(startNumber, endNumber)
		for indexColumn := range ticket.Grid[indexRow] {
			var value int
			value, suffleNumber = suffleNumber[0], suffleNumber[1:]
			ticket.Grid[indexRow][indexColumn].Number = value
		}
	}
	return ticket
}

func (t Ticket) GetGridNumber(Row, Column int) int {
	return t.Grid[Row][Column].Number
}

func (t Ticket) GetGridStatus(Row, Column int) bool {
	return t.Grid[Row][Column].Status
}
