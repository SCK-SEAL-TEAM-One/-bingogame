package bingogame_test

import (
	. "bingogame"
	"testing"
)

<<<<<<< HEAD
func Test_AcceptanceTest_Vertical_CenterBingoRule_Input_Player_A_And_B_PlayRound_Should_Be_Player_B_Bingo_With_Number_45_41_32_36(t *testing.T) {

	ticketBlankPlayerA := NewTicket(5)
	ticketBlankPlayerB := NewTicket(5)
	ticketWithNumberA := MockTicketNumber(ticketBlankPlayerA, 1)
	ticketWithNumberB := MockTicketNumber(ticketBlankPlayerB, 1)
	playerA := NewPlayer("A", ticketWithNumberA)
	playerB := NewPlayer("B", ticketWithNumberB)
	numberBox := NewNumberBox(75)
	allPlayer := [2]Player{playerA, playerB}
	game := NewGame(allPlayer, numberBox)

	bingoPlayer := ""
	for bingoPlayer == "" {

		pickupNumber := game.PickUpNumber()
		positionXPlayer1, positionYPlayer1 := game.Player[0].CheckNumber()

		if positionXPlayer1 != -1 && positionXPlayer1 != -1 {
			game.Player[0].Mark()
		}

		positionXPlayer2, positionYPlayer2 := game.Player[1].CheckNumber()

		if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
			game.Player[1].Mark()
		}

		if game.Player[0].GetBingo() {
			bingoPlayer = game.Player[0].GetName()
			break
		}
		if game.Player[1].GetBingo() {
			bingoPlayer = game.Player[1].GetName()
			break
		}

	}

	expectedBingoPlayer := "B"
	actualBingoPlayer := bingoPlayer

	if expectedBingoPlayer != actualBingoPlayer {
=======
func Test_AcceptanceTest_Vertical_CenterBingoRule_Input_Player_A_And_B_PlayRound_Should_Be_Player_B_Bingo_With_Number_45_41_32_36(t *testing.T){
    
    ticketBlankPlayerA := NewTicket(5)
    ticketBlankPlayerB := NewTicket(5)
    ticketWithNumberA := MockTicketNumber(ticketBlankPlayerA,1)
    ticketWithNumberB := MockTicketNumber(ticketBlankPlayerB,1)
    playerA := NewPlayer("A",ticketWithNumberA)
    playerB := NewPlayer("B",ticketWithNumberB)
    numberBox := NewNumberBox(75)
    allPlayer := []Player{playerA,playerB}
    game := NewGame(allPlayer,numberBox)

    bingoPlayer := ""
    for bingoPlayer == "" {
	
        pickupNumber := game.PickUpNumber()
        positionXPlayer1,positionYPlayer1 := game.Player[0].CheckNumber()

        if positionXPlayer1 != -1 && positionXPlayer1 != -1 {
            game.Player[0].Mark()    
        }

        positionXPlayer2,positionYPlayer2 := game.Player[1].CheckNumber()

        if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
            game.Player[1].Mark()    
        }

        if game.Player[0].GetBingo(){
            bingoPlayer = game.Player[0].GetName()
            break
        }
        if game.Player[1].GetBingo(){
            bingoPlayer = game.Player[1].GetName()
            break
        }

    }

    expectedBingoPlayer := "B"
    actualBingoPlayer := bingoPlayer

    if expectedBingoPlayer != actualBingoPlayer {
>>>>>>> 1732a67e52c7d436f16be5b3931cb7c3545474a0
		t.Errorf("expected player is %s but it got %s", expectedBingoPlayer, actualBingoPlayer)
	}

}
