package bingogame_test

import (
	. "bingogame"
	"testing"
)

func Test_AcceptanceTest_Vertical_CenterBingoRule_Input_Player_A_And_B_PlayRound_8_Should_Be_Player_A_Bingo_With_Number_51_47_56_49_58(t *testing.T) {

	ticketBlankPlayerA := NewTicket(5)
	ticketBlankPlayerB := NewTicket(5)
	ticketWithNumberA := MockTicketNumber(ticketBlankPlayerA, 1)
	ticketWithNumberB := MockTicketNumber(ticketBlankPlayerB, 1)
	playerA := NewPlayer("A", ticketWithNumberA)
	playerB := NewPlayer("B", ticketWithNumberB)
	numberBox := NewNumberBox(75)
	numberBox = MockNumberBox()
	allPlayer := []Player{playerA, playerB}
	game := NewGame(allPlayer, numberBox)

	bingoPlayer := ""
	for bingoPlayer == "" {

		pickupNumber := game.PickUpNumber()
		positionXPlayer1, positionYPlayer1 := game.Players[0].CheckNumber(pickupNumber)

		if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
			game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
			if game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
				bingoPlayer = game.Players[0].Name
				break
			}
		}

		positionXPlayer2, positionYPlayer2 := game.Players[1].CheckNumber(pickupNumber)

		if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
			game.Players[1].Mark(positionXPlayer2, positionYPlayer2)

			if game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
				bingoPlayer = game.Players[1].Name
				break
			}
		}
	}

	expectedBingoPlayer := "B"
	actualBingoPlayer := bingoPlayer

	if expectedBingoPlayer != actualBingoPlayer {
		t.Errorf("expected player is %s but it got %s", expectedBingoPlayer, actualBingoPlayer)
	}

}

func MockNumberBox() []int {
	return []int{9, 51, 47, 29, 56, 49, 39, 58}
}
