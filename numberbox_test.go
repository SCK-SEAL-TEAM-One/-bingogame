package bingogame_test

import (
	. "bingogame"
	"testing"
)

func Test_NewNumberBox_Input_Number_75_Should_Be_NumberBox_Have_Array_1_To_75_In_NumberBox(t *testing.T) {
	Number := 75
	var expectedNumberBox [75]int
	for i := 0; i < Number; i++ {
		expectedNumberBox[i] = i + 1
	}

	actual := NewNumberBox(Number)

	if expectedNumberBox != actual {
		t.Errorf("expectedNumberbox %v but got %v", expectedNumberBox, actual)
	}

