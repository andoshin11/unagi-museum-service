package repository

import (
	"context"
	"log"

	"github.com/andoshin11/euphro-ddd/src/entity"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

// MuseumRepository interface
type MuseumRepository interface {
	GetAll(ctx context.Context) ([]*entity.Museum, error)
	GetNeighborsByLat(ctx context.Context, lat float64, distance int) ([]*entity.Museum, error)
	GetByID(ctx context.Context, id string) (*entity.Museum, error)
}

type museumRepository struct {
	Client *firestore.Client
}

// NewMuseumRepository return struct
func NewMuseumRepository(Client *firestore.Client) MuseumRepository {
	return &museumRepository{Client}
}

func (r *museumRepository) GetAll(ctx context.Context) ([]*entity.Museum, error) {
	museums := []*entity.Museum{}

	iter := r.Client.Collection("museum").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		museum := entity.Museum{}
		doc.DataTo(&museum)
		museums = append(museums, &museum)
	}

	return museums, nil
}

func (r *museumRepository) GetNeighborsByLat(ctx context.Context, lat float64, distance int) ([]*entity.Museum, error) {
	distInFloat64 := float64(distance)
	distInLat := distInFloat64 * 0.0090133729745762
	latStart := lat - distInLat
	latEnd := lat + distInLat

	museums := []*entity.Museum{}

	iter := r.Client.Collection("museum").Where("lat", ">", latStart).Where("lat", "<", latEnd).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		museum := entity.Museum{}
		doc.DataTo(&museum)
		museums = append(museums, &museum)
	}

	return museums, nil
}

func (r *museumRepository) GetByID(ctx context.Context, id string) (*entity.Museum, error) {

	snapshot, err := r.Client.Collection("museum").Doc(id).Get(ctx)
	museum := entity.Museum{}
	snapshot.DataTo(&museum)
	return &museum, err
}
