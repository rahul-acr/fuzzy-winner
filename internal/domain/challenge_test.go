package domain

import (
	"testing"
	"time"
)

func TestRahulShouldBeAbleToChallengeParikshit(t *testing.T) {
	parikshit := &Player{id: 1}
	rahul := &Player{id: 2}

	challenge := rahul.challenge(parikshit)

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

	challenge := rahul.challenge(parikshit)

	datetime, err := time.Parse(time.RFC3339, "2023-01-02T15:04:00+05:30")
	if err != nil {
		t.Fatal(err)
	}

	parikshit.accept(challenge, datetime)

	if !challenge.isAccepted {
		t.Fatalf("Challenge should be accepted")
	}
	if challenge.time != datetime {
		t.Fatalf("date time does not match")
	}

}
