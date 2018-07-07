package usecase

import (
	"context"
	"log"

	"github.com/andoshin11/euphro-ddd/src/entity"

	"github.com/andoshin11/euphro-ddd/src/repository"
)

// MuseumUsecase type
type MuseumUsecase interface {
	GetAll(ctx context.Context) ([]*entity.Museum, error)
	GetNeighbors(ctx context.Context, lat float64, lng float64, distance int) ([]*entity.Museum, error)
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

func (u *museumUsecase) GetNeighbors(ctx context.Context, lat float64, lng float64, distance int) ([]*entity.Museum, error) {
	neighborsByLat, err := u.MuseumRepository.GetNeighborsByLat(ctx, lat, distance)

	if err != nil {
		log.Fatalln(err)
	}

	distInFloat64 := float64(distance)
	distInLng := distInFloat64 * 0.010966404715491394
	lngStart := lng - distInLng
	lngEnd := lng + distInLng

	neighborsByLng := neighborsByLat[:0]

	for _, x := range neighborsByLat {
		targetLng := x.Longitude

		if lngStart < targetLng && targetLng < lngEnd {
			neighborsByLng = append(neighborsByLng, x)
		}
	}

	return neighborsByLng, nil
}
