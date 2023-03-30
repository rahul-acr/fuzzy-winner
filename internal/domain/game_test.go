package domain

import "testing"

func TestWinsShouldIncrementWhenPlayerWinsAMatch(t *testing.T) {
	parikshit := &Player{1, 0, 0}
	rahul := &Player{2, 0, 0}
	match := Between(parikshit, rahul)
	match.wonBy(parikshit)
	if parikshit.wins != 1 {
		t.Fatalf("Parikshit should have 1 win")
	}
}

func TestLossesShouldIncrementWhenPlayerLosesAMatch(t *testing.T) {
	parikshit := &Player{1, 0, 0}
	rahul := &Player{2, 0, 0}
	match := Between(parikshit, rahul)
	match.wonBy(rahul)
	if parikshit.losses != 1 {
		t.Fatalf("Parikshit should have 1 loss")
	}
}
