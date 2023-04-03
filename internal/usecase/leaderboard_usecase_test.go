package usecase

import (
	"testing"
	"tv/quick-bat/internal/domain"
)

func TestParikshitShouldHave1WinAndRahulHave1LossWhenParikshitAddsAWinMatchAgainstRahul(t *testing.T) {
	parikshit := domain.CreatePlayer(1, 0, 0)
	rahul := domain.CreatePlayer(2, 0, 0)

	domain.TtLeaderBoard.Init([]*domain.Player{parikshit, rahul})
	match := Match{1, 2, true}
	AddMatch(match)

	if parikshit.Wins() != 1 || rahul.Losses() != 1 {
		t.Fatalf("Parikshit should have 1 win and Rahul have 1 loss")
	}
}

func TestRahulShouldHave1WinAndParikshitHave1LossWhenParikshitAddsALoseMatchAgainstRahul(t *testing.T) {
	parikshit := domain.CreatePlayer(1, 0, 0)
	rahul := domain.CreatePlayer(2, 0, 0)

	domain.TtLeaderBoard.Init([]*domain.Player{parikshit, rahul})
	match := Match{1, 2, false}
	AddMatch(match)

	if parikshit.Losses() != 1 || rahul.Wins() != 1 {
		t.Fatalf("Parikshit should have 1 loss and Rahul have 1 win")
	}
}
