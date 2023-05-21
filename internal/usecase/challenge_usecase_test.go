package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"tv/quick-bat/internal/domain"
)

func TestChallengeShouldBeCreatedWhenRahulChallengesParikshit(t *testing.T) {
	rahul := domain.NewPlayer(1, 1, 1)
	parikshit := domain.NewPlayer(2, 1, 1)
	domain.MainLeaderBoard = domain.NewLeaderBoard([]*domain.Player{parikshit, rahul})

	var challenge *domain.Challenge
	domain.OnChallengeCreate = func(c *domain.Challenge) {
		challenge = c
	}

	CreateChallenge(Challenge{
		Challenger: int(rahul.Id()),
		Opponent:   int(parikshit.Id()),
	})

	a := assert.New(t)
	a.NotNil(challenge)
	a.Equal(rahul, challenge.Challenger())
	a.Equal(parikshit, challenge.Opponent())
	a.Nil(challenge.Time())
	a.False(challenge.IsAccepted())
}

func TestParikshitShouldBeAbleToAcceptTheChallenge(t *testing.T) {
	rahul := domain.NewPlayer(1, 1, 1)
	parikshit := domain.NewPlayer(2, 1, 1)
	domain.MainLeaderBoard = domain.NewLeaderBoard([]*domain.Player{parikshit, rahul})

	const createdChallengeId = 1
	var challenge *domain.Challenge
	domain.OnChallengeCreate = func(c *domain.Challenge) {
		c.Id = createdChallengeId
		challenge = c
	}

	domain.LoadChallenge = func(challengeId interface{}) *domain.Challenge {
		if challengeId == createdChallengeId {
			return challenge
		}
		return nil
	}

	CreateChallenge(Challenge{
		Challenger: int(rahul.Id()),
		Opponent:   int(parikshit.Id()),
	})

	matchTime := time.Now().Add(time.Hour * 2)
	AcceptChallenge(createdChallengeId, int(parikshit.Id()), matchTime)

	a := assert.New(t)
	a.NotNil(challenge)
	a.True(challenge.IsAccepted())
	a.Equal(matchTime, *challenge.Time())
}
