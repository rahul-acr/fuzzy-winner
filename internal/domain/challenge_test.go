package domain

import (
	"testing"
	"time"
)

func TestRahulShouldBeAbleToChallengeParikshit(t *testing.T) {
	parikshit := Player{id: 1}
	rahul := Player{id: 2}

	challenge := rahul.Challenge(parikshit)

	if challenge.challenger != rahul {
		t.Fatalf("Rahul should be the challenger")
	}
	if challenge.opponent != parikshit {
		t.Fatalf("Parikshit should be the opponent")
	}
}

func TestParikshitShouldBeAbleToAcceptChallengeFromRahul(t *testing.T) {
	parikshit := Player{id: 1}
	rahul := Player{id: 2}

	matchTime := time.Now().Add(time.Hour * 2)
	false
	challenge := rahul.Challenge(parikshit)
	parikshit.Accept(&challenge, matchTime)

	if !challenge.isAccepted {
		t.Fatalf("Challenge should be accepted")
	}
	if *challenge.matchTime != matchTime {
		t.Fatalf("date time does not match")
	}

}

func TestLeaderBoardShouldUpdateWhenParikshitWon(t *testing.T) {
	parikshit := Player{id: 1}
	rahul := Player{id: 2}
	MainLeaderBoard = NewLeaderBoard([]Player{parikshit, rahul})

	matchTime := time.Now().Add(time.Hour * 2)
	challenge := rahul.Challenge(parikshit)
	parikshit.Accept(&challenge, matchTime)

	challenge.winBy(parikshit)

	if MainLeaderBoard.GetRank(parikshit) != 1 {
		t.Fatalf("Parikshit's rank should be 1")
	}
	if MainLeaderBoard.GetRank(rahul) != 2 {
		t.Fatalf("Rahul's rank should be 2")
	}
}
