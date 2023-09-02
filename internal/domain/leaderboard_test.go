package domain

import (
	"testing"
	"tv/quick-bat/internal/events"
)

func TestParikshitShouldHaveRank1AndRahulHaveRank2WhenParikshitWinsAMatchAgainstRahul(t *testing.T) {
	parikshit := Player{id: 1}
	rahul := Player{id: 2}

	MainLeaderBoard = NewLeaderBoard([]Player{parikshit, rahul})
	defer events.Clear("playerUpdate")

	parikshit.WinAgainst(&rahul)

	if MainLeaderBoard.GetRank(parikshit) != 1 {
		t.Fatalf("Parikshit's rank should be 1")
	}
	if MainLeaderBoard.GetRank(rahul) != 2 {
		t.Fatalf("Rahul's rank should be 2")
	}
}

func TestRahulShouldHaveRank1WhenHeOvertakesParikshitInWins(t *testing.T) {
	parikshit := Player{id: 1}
	rahul := Player{id: 2}

	MainLeaderBoard = NewLeaderBoard([]Player{parikshit, rahul})
	defer events.Clear("playerUpdate")

	parikshit.WinAgainst(&rahul)
	rahul.WinAgainst(&parikshit)
	rahul.WinAgainst(&parikshit)

	if MainLeaderBoard.GetRank(rahul) != 1 {
		t.Fatalf("Rahul's rank should be 1")
	}
}

func TestParikshitsRankShouldBe2WhenHarunScoresMoreWinsThanHim(t *testing.T) {
	parikshit := Player{id: 1}
	rahul := Player{id: 2}
	harun := Player{id: 3}

	MainLeaderBoard = NewLeaderBoard([]Player{parikshit, rahul, harun})
	defer events.Clear("playerUpdate")

	parikshit.WinAgainst(&rahul)
	harun.WinAgainst(&parikshit)
	harun.WinAgainst(&rahul)

	if MainLeaderBoard.GetRank(parikshit) != 2 {
		t.Fatalf("Parikshit's rank should be 2")
	}
}

func TestLeaderBoardReuseOldMatchData(t *testing.T) {
	parikshit := Player{id: 1, wins: 2}
	rahul := Player{id: 2, wins: 3}

	MainLeaderBoard = NewLeaderBoard([]Player{parikshit, rahul})
	defer events.Clear("playerUpdate")

	if MainLeaderBoard.GetRank(parikshit) != 2 {
		t.Fatalf("Parikshit's rank should be 2")
	}
	if MainLeaderBoard.GetRank(rahul) != 1 {
		t.Fatalf("Rahul's rank should be 1")
	}
}
