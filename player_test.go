package bingogame_test

import(
    . "bingogame"
    "testing"
)

func Test_NewPlayer_Input_A_Should_Be_Struct(t *testing.T) {
    playerName := "A"
    expected := Player{Name:"A",Ticket: Ticket{}}

    actual := NewPlayer(playerName)
    
    if expected != actual{
        t.Errorf("Should be %v but got %v",expected,actual)
    }

}