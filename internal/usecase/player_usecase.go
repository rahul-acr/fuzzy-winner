package usecase

import (
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
)

type PlayerManager struct {
	PlayerRepository db.PlayerRepository
}

func (p *PlayerManager) FindPlayer(playerId int) (domain.Player, error) {
	playerRecord, err := p.PlayerRepository.FindPlayer(playerId)
	if err != nil {
		return domain.Player{}, err
	}
	return domain.NewPlayer2(domain.PlayerId(playerId), playerRecord.Wins, playerRecord.Losses), nil

}
