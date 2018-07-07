package usecase

import (
	"context"

	"github.com/andoshin11/euphro-ddd/src/entity"

	"github.com/andoshin11/euphro-ddd/src/repository"
)

// MuseumUsecase type
type MuseumUsecase interface {
	GetAll(ctx context.Context) ([]*entity.Museum, error)
}

type museumUsecase struct {
	MuseumRepository repository.MuseumRepository
}

// NewMuseumUsecase return a ref
func NewMuseumUsecase(repo repository.MuseumRepository) MuseumUsecase {
	return &museumUsecase{
		MuseumRepository: repo,
	}
}

func (u *museumUsecase) GetAll(ctx context.Context) ([]*entity.Museum, error) {
	return u.MuseumRepository.GetAll(ctx)
}
