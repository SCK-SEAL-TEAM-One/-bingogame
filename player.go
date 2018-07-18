package bingogame

import (
	"fmt"
)

func NewPlayer(name string, ticket Ticket) Player {
	return Player{
		Name:   name,
		Ticket: ticket,
	}
}

func (p *Player) Mark(positionX, positionY int) bool {
	positionX -= 1
	positionY -= 1
	p.Ticket.Grid[positionX][positionY].Status = true

	if p.Ticket.Grid[positionX][positionY].Status == false {
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
	fmt.Println(positionX, positionY)
	return p.CheckHorizontal(positionX, positionY) || p.CheckDiagonal(positionX, positionY) || p.CheckVertical(positionX, positionY)
}

func (p Player) CheckVertical(positionX, positionY int) bool {
	positionX--
	positionY--
	var number int
	for rowIndex := range p.Ticket.Grid {
		if p.Ticket.Grid[rowIndex][positionY].Status == true {
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
