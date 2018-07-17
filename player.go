package bingogame

func NewPlayer(name string, ticket Ticket) Player {
	return Player{
		Name:   name,
		Ticket: ticket,
	}
}

func (p *Player)Mark(x,y int) bool {
	p.Ticket.Grid[x][y].Status = true
	return true

	if p.Ticket.Grid[x][y].Status == false {
		return false
	}
	
}