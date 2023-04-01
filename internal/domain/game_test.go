package domain

import "testing"

func TestParikshitShouldHave1WinsWhenHeWinsAMatch(t *testing.T) {
	parikshit := &Player{1, 0, 0}
	rahul := &Player{2, 0, 0}
	parikshit.WinAgainst(rahul)
	if parikshit.wins != 1 {
		t.Fatalf("Parikshit should have 1 win")
	}
}

func TestParikshitShouldHave1LossWhenHeLosesAMatch(t *testing.T) {
	parikshit := &Player{1, 0, 0}
	rahul := &Player{2, 0, 0}
	rahul.WinAgainst(parikshit)
	if parikshit.losses != 1 {
		t.Fatalf("Parikshit should have 1 loss")
	}
}

func TestRahulShouldHave2WinsWhenHeWinsTwoMatches(t *testing.T) {
	parikshit := &Player{1, 0, 0}
	rahul := &Player{2, 0, 0}
	rahul.WinAgainst(parikshit)
	rahul.WinAgainst(parikshit)
	if rahul.wins != 2 {
		t.Fatalf("Rahul should have 2 wins")
	}
}

func TestRahulShouldHave2LossessWhenHeLosesTwoMatches(t *testing.T) {
	parikshit := &Player{1, 0, 0}
	rahul := &Player{2, 0, 0}
	parikshit.WinAgainst(rahul)
	parikshit.WinAgainst(rahul)
	if parikshit.wins != 2 {
		t.Fatalf("Parikshit should have 2 wins")
	}
}
