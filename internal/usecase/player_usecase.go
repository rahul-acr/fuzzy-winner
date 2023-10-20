package usecase

import (
	"context"
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
)

type PlayerManager struct {
	PlayerRepository db.PlayerRepository
}

type PlayerInfoShort struct {
	PlayerId any    `json:"playerId"`
	Name     string `json:"name"`
}

func (p PlayerManager) FindPlayer(ctx context.Context, playerId int) (domain.Player, error) {
	return p.PlayerRepository.FindPlayer(ctx, playerId)
}

func (p PlayerManager) FindPlayers(ctx context.Context) []domain.Player {
	return p.PlayerRepository.FetchAll(ctx)
}
