package domain

import "testing"

func TestShouldMakeWinnerRank1AndLoserRank2IfItIsTheFirstMatch(t *testing.T) {
	parikshit := Player{1, 0, 0}
	rahul := Player{2, 0, 0}

	leaderBoard.Init([]Player{parikshit, rahul})

	parikshit.WinAgainst(&rahul)

	if parikshit.GetRank() != 1 {
		t.Fatalf("Parikshit's rank should be 1")
	}
	if rahul.GetRank() != 2 {
		t.Fatalf("Rahul's rank should be 2")
	}
}
