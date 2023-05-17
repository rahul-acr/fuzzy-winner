package domain

import (
	"testing"
	"time"
)

func TestRahulShouldBeAbleToChallengeParikshit(t *testing.T) {
	parikshit := &Player{id: 1}
	rahul := &Player{id: 2}

	challenge := rahul.Challenge(parikshit)

	if challenge == nil {
		t.Fatalf("Challenge should be added")
	}
	if challenge.challenger != rahul {
		t.Fatalf("Rahul should be the challenger")
	}
	if challenge.opponent != parikshit {
		t.Fatalf("Parikshit should be the opponent")
	}
}

func TestParikshitShouldBeAbleToAcceptChallengeFromRahul(t *testing.T) {
	parikshit := &Player{id: 1}
	rahul := &Player{id: 2}

	matchTime := time.Now().Add(time.Hour * 2)
	challenge := rahul.Challenge(parikshit)
	parikshit.Accept(challenge, matchTime)

	if !challenge.isAccepted {
		t.Fatalf("Challenge should be accepted")
	}
	if challenge.time != matchTime {
		t.Fatalf("date time does not match")
	}

}

func TestLeaderBoardShouldUpdateWhenParikshitWon(t *testing.T) {
	parikshit := &Player{id: 1}
	rahul := &Player{id: 2}
	leaderBoard.Init([]*Player{parikshit, rahul})

	matchTime := time.Now().Add(time.Hour * 2)
	challenge := rahul.Challenge(parikshit)
	parikshit.Accept(challenge, matchTime)

	challenge.WonBy(parikshit)

	if leaderBoard.GetRank(parikshit) != 1 {
		t.Fatalf("Parikshit's rank should be 1")
	}
	if leaderBoard.GetRank(rahul) != 2 {
		t.Fatalf("Rahul's rank should be 2")
	}
}
