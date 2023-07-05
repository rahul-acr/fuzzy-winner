package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"tv/quick-bat/internal/domain"
	"tv/quick-bat/internal/events"
)

func TestChallengeShouldBeCreatedWhenRahulChallengesParikshit(t *testing.T) {
	rahul := domain.NewPlayer(1, 1, 1)
	parikshit := domain.NewPlayer(2, 1, 1)
	domain.MainLeaderBoard = domain.NewLeaderBoard([]*domain.Player{parikshit, rahul})

	var challenge *domain.Challenge

	events.Listen("challengeCreate", func(event events.Event) {
		challenge = event.Payload.(*domain.Challenge)
	})

	CreateChallenge(Challenge{
		ChallengerId: int(rahul.Id()),
		OpponentId:   int(parikshit.Id()),
	})

	a := assert.New(t)
	a.NotNil(challenge)
	a.Equal(rahul, challenge.Challenger())
	a.Equal(parikshit, challenge.Opponent())
	a.Nil(challenge.Time())
	a.False(challenge.IsAccepted())
}

func TestParikshitShouldBeAbleToAcceptTheChallenge(t *testing.T) {
	a := assert.New(t)

	rahul := domain.NewPlayer(1, 1, 1)
	parikshit := domain.NewPlayer(2, 1, 1)
	domain.MainLeaderBoard = domain.NewLeaderBoard([]*domain.Player{parikshit, rahul})

	const createdChallengeId = 1
	var challenge *domain.Challenge

	events.Listen("challengeCreate", func(event events.Event) {
		challenge = event.Payload.(*domain.Challenge)
		challenge.Id = createdChallengeId
	})

	LoadChallenge = func(challengeId interface{}) (*domain.Challenge, error) {
		if challengeId == createdChallengeId {
			return challenge, nil
		}
		return nil, errors.New("unknown challenge id")
	}

	CreateChallenge(Challenge{
		ChallengerId: int(rahul.Id()),
		OpponentId:   int(parikshit.Id()),
	})

	matchTime, err := time.Parse(time.RFC3339, "2023-07-05T10:45:26Z")
	a.Nil(err)

	err = AcceptChallenge(createdChallengeId, ChallengeAccept{
		int(parikshit.Id()),
		matchTime,
	})

	a.Nil(err)
	a.NotNil(challenge)
	a.True(challenge.IsAccepted())
	a.Equal(matchTime, *challenge.Time())
}
