package usecase

import (
	"context"
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
)

type PlayerManager struct {
	PlayerRepository db.PlayerRepository
}

func (p *PlayerManager) FindPlayer(ctx context.Context, playerId int) (domain.Player, error) {
	playerRecord, err := p.PlayerRepository.FindPlayer(ctx, playerId)
	if err != nil {
		return domain.Player{}, err
	}
	return domain.NewPlayer(domain.PlayerId(playerId), playerRecord.Wins, playerRecord.Losses), nil
}
