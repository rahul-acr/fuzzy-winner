package usecase

import (
	"context"
	"testing"
	"tv/quick-bat/internal/domain"
)

func TestParikshitShouldHave1WinAndRahulHave1LossWhenParikshitAddsAWinMatchAgainstRahul(t *testing.T) {
	var parikshitId domain.PlayerId = 1
	var rahulId domain.PlayerId = 2
	domain.MainLeaderBoard = domain.NewLeaderBoard([]domain.Player{
		domain.NewPlayer(parikshitId, 0, 0),
		domain.NewPlayer(rahulId, 0, 0),
	})
	match := Match{1, 2, true}
	err := AddMatch(context.TODO(), &match)
	if err != nil {
		t.Fatal(err)
	}
	parikshit, err := domain.MainLeaderBoard.FindPlayer(parikshitId)
	if err != nil {
		t.Fatal(err)
	}
	rahul, err := domain.MainLeaderBoard.FindPlayer(rahulId)
	if err != nil {
		t.Fatal(err)
	}
	if parikshit.Wins() != 1 || rahul.Losses() != 1 {
		t.Fatalf("Parikshit should have 1 win and Rahul have 1 loss")
	}
}

func TestRahulShouldHave1WinAndParikshitHave1LossWhenParikshitAddsALoseMatchAgainstRahul(t *testing.T) {
	var parikshitId domain.PlayerId = 1
	var rahulId domain.PlayerId = 2
	domain.MainLeaderBoard = domain.NewLeaderBoard([]domain.Player{
		domain.NewPlayer(parikshitId, 0, 0),
		domain.NewPlayer(rahulId, 0, 0),
	})
	match := Match{1, 2, false}
	err := AddMatch(context.TODO(), &match)
	if err != nil {
		t.Fatal(err)
	}
	parikshit, err := domain.MainLeaderBoard.FindPlayer(parikshitId)
	if err != nil {
		t.Fatal(err)
	}
	rahul, err := domain.MainLeaderBoard.FindPlayer(rahulId)
	if err != nil {
		t.Fatal(err)
	}

	if parikshit.Losses() != 1 || rahul.Wins() != 1 {
		t.Fatalf("Parikshit should have 1 loss and Rahul have 1 win")
	}
}

func TestShouldGivePlayerDetails(t *testing.T) {
	parikshit := domain.NewPlayer(1, 0, 0)
	rahul := domain.NewPlayer(2, 0, 0)

	domain.MainLeaderBoard = domain.NewLeaderBoard([]domain.Player{parikshit, rahul})

	rahul.WinAgainst(&parikshit)
	parikshit.WinAgainst(&rahul)
	rahul.WinAgainst(&parikshit)

	rahulsDetails, _ := GetPlayerDetails(context.TODO(), 2)

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
