package bingogame_test

import (
	. "bingogame"
	"testing"
)

func Test_NewNumberBox_Input_75_Should_Be_Number_1_To_75(t *testing.T) {
	number := 75
	expectedlength := 75
	actualNumbers := NewNumberBox(number)

	if expectedlength != len(actualNumbers) {
		t.Errorf("expected Array Length  %v but got %v", expectedlength, len(actualNumbers))
	}

	unique := map[int]bool{}
	for index := 0; index < number; index++ {
		value := actualNumbers[index]
		if unique[value] {
			t.Errorf("expected Array is Unique")
			break
		}
		unique[value] = true
	}
}
