package bingogame

type State struct{
    Status bool
    Number int
}

type Ticket struct{
    Grid [][]State
}

type Player struct{
    Name string
    Ticket Ticket
}

func NewPlayer(playerName string) Player{
	return Player{Name: playerName,Ticket: Ticket{}}
}