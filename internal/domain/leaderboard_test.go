package domain

import (
	"testing"
)

func TestParikshitShouldHaveRank1AndRahulHaveRank2WhenParikshitWinsAMatchAgainstRahul(t *testing.T) {
	parikshit := &Player{1, 0, 0}
	rahul := &Player{2, 0, 0}

	leaderBoard.Init([]*Player{parikshit, rahul})

	parikshit.WinAgainst(rahul)

	if parikshit.GetRank() != 1 {
		t.Fatalf("Parikshit's rank should be 1")
	}
	if rahul.GetRank() != 2 {
		t.Fatalf("Rahul's rank should be 2")
	}
}

func TestRahulShouldHaveRank1WhenHeOvertakesParikshitInWins(t *testing.T) {
	parikshit := &Player{1, 0, 0}
	rahul := &Player{2, 0, 0}

	leaderBoard.Init([]*Player{parikshit, rahul})

	parikshit.WinAgainst(rahul)
	rahul.WinAgainst(parikshit)
	rahul.WinAgainst(parikshit)

	if rahul.GetRank() != 1 {
		t.Fatalf("Rahul's rank should be 1")
	}
}

func TestParikshitsRankShouldBe2WhenHarunScoresMoreWinsThanHim(t *testing.T) {
	parikshit := &Player{1, 0, 0}
	rahul := &Player{2, 0, 0}
	harun := &Player{3, 0, 0}

	leaderBoard.Init([]*Player{parikshit, rahul, harun})

	parikshit.WinAgainst(rahul)
	harun.WinAgainst(parikshit)
	harun.WinAgainst(rahul)

	if parikshit.GetRank() != 2 {
		t.Fatalf("Parikshit's rank should be 2")
	}
}
