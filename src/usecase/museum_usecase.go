package usecase

import (
	"context"
	"log"

	"github.com/andoshin11/unagi-museum-service/src/entity"

	"github.com/andoshin11/unagi-museum-service/src/repository"
)

// MuseumUsecase type
type MuseumUsecase interface {
	GetAll(ctx context.Context) ([]*entity.Museum, error)
	GetNeighbors(ctx context.Context, lat float64, lng float64, distance int) ([]*entity.Museum, error)
	GetByID(ctx context.Context, id string) (*entity.Museum, error)
}

type museumUsecase struct {
	museumRepository repository.MuseumRepository
}

// NewMuseumUsecase return a ref
func NewMuseumUsecase(repo repository.MuseumRepository) MuseumUsecase {
	return &museumUsecase{
		museumRepository: repo,
	}
}

func (u *museumUsecase) GetAll(ctx context.Context) ([]*entity.Museum, error) {
	return u.museumRepository.GetAll(ctx)
}

func (u *museumUsecase) GetNeighbors(ctx context.Context, lat float64, lng float64, distance int) ([]*entity.Museum, error) {
	neighborsByLat, err := u.museumRepository.GetNeighborsByLat(ctx, lat, distance)

	if err != nil {
		log.Fatalln(err)
	}

	distInFloat64 := float64(distance)
	distInLng := distInFloat64 * 0.010966404715491394
	lngStart := lng - distInLng
	lngEnd := lng + distInLng

	neighborsByLng := neighborsByLat[:0]

	for _, x := range neighborsByLat {
		targetLng := x.Lng

		if lngStart < targetLng && targetLng < lngEnd {
			neighborsByLng = append(neighborsByLng, x)
		}
	}

	return neighborsByLng, nil
}

func (u *museumUsecase) GetByID(ctx context.Context, id string) (*entity.Museum, error) {
	return u.museumRepository.GetByID(ctx, id)
}
