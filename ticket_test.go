package bingogame_test

import (
	. "bingogame"
	"fmt"
	"testing"
)

func Test_NewTicket_Input_Number_Of_Grid_5_Should_Be_Ticket_With_SizeX_And_SizeY_And_Grid(t *testing.T) {
	numberOfGrid := 5
	grid := make([][]State, numberOfGrid)
	for index := 0; index < numberOfGrid; index++ {
		grid[index] = make([]State, numberOfGrid)
	}
	expected := Ticket{SizeX: 5, SizeY: 5, Grid: grid}

	actualTicket := NewTicket(numberOfGrid)

	for indexRow := 0; indexRow < 5; indexRow++ {
		for indexColumn := 0; indexColumn < 5; indexColumn++ {
			if expected.Grid[indexRow][indexColumn] != actualTicket.Grid[indexRow][indexColumn] {
				t.Errorf("expect %v but got %v", expected, actualTicket)
			}
		}
	}

	if actualTicket.SizeX != expected.SizeX || actualTicket.SizeY != expected.SizeY {
		t.Errorf("expect %v but got %v", expected, actualTicket)
	}
}

func Test_GenerateTicketNumber_Input_BlankTicket_Should_Be_TicketWithNumber(t *testing.T) {
	grid := make([][]State, 5)
	for index := 0; index < 5; index++ {
		grid[index] = make([]State, 5)
	}
	ticket := Ticket{SizeX: 5, SizeY: 5, Grid: grid}

	ticket = GenerateTicketNumber(ticket)
	fmt.Printf("!%v", ticket.Grid[0][0])
	for indexRow := 0; indexRow < 5; indexRow++ {
		for indexColumn := 0; indexColumn < 5; indexColumn++ {
			if ticket.Grid[indexRow][indexColumn].Number == 0 && indexColumn != 2 && indexColumn != 2 {
				t.Errorf("GirdNumber Row %d Column %d is %d", indexRow, indexColumn, ticket.Grid[indexRow][indexColumn].Number)
			}
		}
	}
}
