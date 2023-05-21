package usecase

import (
	"testing"
	"tv/quick-bat/internal/domain"
)

func TestParikshitShouldHave1WinAndRahulHave1LossWhenParikshitAddsAWinMatchAgainstRahul(t *testing.T) {
	parikshit := domain.NewPlayer(1, 0, 0)
	rahul := domain.NewPlayer(2, 0, 0)

	domain.MainLeaderBoard = domain.NewLeaderBoard([]*domain.Player{parikshit, rahul})
	match := Match{1, 2, true}
	AddMatch(&match)

	if parikshit.Wins() != 1 || rahul.Losses() != 1 {
		t.Fatalf("Parikshit should have 1 win and Rahul have 1 loss")
	}
}

func TestRahulShouldHave1WinAndParikshitHave1LossWhenParikshitAddsALoseMatchAgainstRahul(t *testing.T) {
	parikshit := domain.NewPlayer(1, 0, 0)
	rahul := domain.NewPlayer(2, 0, 0)

	domain.MainLeaderBoard = domain.NewLeaderBoard([]*domain.Player{parikshit, rahul})
	match := Match{1, 2, false}
	AddMatch(&match)

	if parikshit.Losses() != 1 || rahul.Wins() != 1 {
		t.Fatalf("Parikshit should have 1 loss and Rahul have 1 win")
	}
}

func TestShouldGivePlayerDetails(t *testing.T) {
	parikshit := domain.NewPlayer(1, 0, 0)
	rahul := domain.NewPlayer(2, 0, 0)

	domain.MainLeaderBoard = domain.NewLeaderBoard([]*domain.Player{parikshit, rahul})

	rahul.WinAgainst(parikshit)
	parikshit.WinAgainst(rahul)
	rahul.WinAgainst(parikshit)

	rahulsDetails := GetPlayerDetails(2)

	if rahulsDetails.Wins != 2 {
		t.Fatalf("Rahul should have 2 wins")
	}
	if rahulsDetails.Losses != 1 {
		t.Fatalf("Rahul should have 1 loss")
	}
	if rahulsDetails.Rank != 1 {
		t.Fatalf("Rahul should have rank 1")
	}
	if rahulsDetails.Id != 2 {
		t.Fatalf("Rahul's playerId should be 2")
	}
}
