package domain

import (
	"testing"
)

func TestParikshitShouldHaveRank1AndRahulHaveRank2WhenParikshitWinsAMatchAgainstRahul(t *testing.T) {
	parikshit := &Player{id: 1}
	rahul := &Player{id: 2}

	TtLeaderBoard.Init([]*Player{parikshit, rahul})

	parikshit.WinAgainst(rahul)

	if TtLeaderBoard.GetRank(parikshit) != 1 {
		t.Fatalf("Parikshit's rank should be 1")
	}
	if TtLeaderBoard.GetRank(rahul) != 2 {
		t.Fatalf("Rahul's rank should be 2")
	}
}

func TestRahulShouldHaveRank1WhenHeOvertakesParikshitInWins(t *testing.T) {
	parikshit := &Player{id: 1}
	rahul := &Player{id: 2}

	TtLeaderBoard.Init([]*Player{parikshit, rahul})

	parikshit.WinAgainst(rahul)
	rahul.WinAgainst(parikshit)
	rahul.WinAgainst(parikshit)

	if TtLeaderBoard.GetRank(rahul) != 1 {
		t.Fatalf("Rahul's rank should be 1")
	}
}

func TestParikshitsRankShouldBe2WhenHarunScoresMoreWinsThanHim(t *testing.T) {
	parikshit := &Player{id: 1}
	rahul := &Player{id: 2}
	harun := &Player{id: 3}

	TtLeaderBoard.Init([]*Player{parikshit, rahul, harun})

	parikshit.WinAgainst(rahul)
	harun.WinAgainst(parikshit)
	harun.WinAgainst(rahul)

	if TtLeaderBoard.GetRank(parikshit) != 2 {
		t.Fatalf("Parikshit's rank should be 2")
	}
}

func TestLeaderBoardReuseOldMatchData(t *testing.T) {
	parikshit := &Player{id: 1, wins: 2}
	rahul := &Player{id: 2, wins: 3}

	TtLeaderBoard.Init([]*Player{parikshit, rahul})

	if TtLeaderBoard.GetRank(parikshit) != 2 {
		t.Fatalf("Parikshit's rank should be 2")
	}
	if TtLeaderBoard.GetRank(rahul) != 1 {
		t.Fatalf("Rahul's rank should be 1")
	}
}
