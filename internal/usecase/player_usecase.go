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
	return p.PlayerRepository.FindPlayer(ctx, playerId)
}
