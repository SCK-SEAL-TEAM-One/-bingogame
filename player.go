package bingogame

func NewPlayer(name string, ticket Ticket) Player {
	return Player{
		Name:   name,
		Ticket: ticket,
	}
}
