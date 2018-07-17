package bingogame

func NewPlayer(name string, ticket Ticket) Player {
	return Player{
		Name:   name,
		Ticket: ticket,
	}
}

func (p *Player)Mark(x,y int) bool {
	p.Ticket.Grid[x][y].Status = true

	if p.Ticket.Grid[x][y].Status == false {
		return false
	}
	return true
}

func (p Player) CheckNumber(number int) (int, int) {
	for rowIndex := range p.Ticket.Grid {
		for columnIndex := range p.Ticket.Grid[rowIndex] {
			if number == p.Ticket.Grid[rowIndex][columnIndex].Number {
				return rowIndex + 1, columnIndex + 1
			}
		}
	}
	return -1, -1
}

func (p Player) GetBingo(positionX, positionY int) bool {
	return p.CheckHorizontal(positionX, positionY) || p.CheckDiagonal(positionX, positionY) || p.CheckVertical(positionX, positionY)
}

func (p Player) CheckVertical(positionX, positionY int) bool {
	var number int
	for rowIndex := range p.Ticket.Grid[positionX] {

		if p.Ticket.Grid[positionX][rowIndex].Status == true {
			number++
			if number == p.Ticket.SizeX {
				return true
			}
		}

	}
	return false
}

func (p Player) CheckDiagonal(positionX, positionY int) bool {
	return false
}

func (p Player) CheckHorizontal(positionX, positionY int) bool {
	return false
}
