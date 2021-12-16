package ports

import "github.com/Haleluak/hexagonal/internal/core/domain"

type GamesRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}