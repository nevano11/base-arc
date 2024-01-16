package repository

import (
	"awesomeProject/internal/entity"
	"context"
	"errors"
	"slices"
)

type FakeBandRepository struct {
	bandArray []entity.Band
}

func NewFakeBandRepository(ctx context.Context) (*FakeBandRepository, error) {
	bands := make([]entity.Band, 3, 3)
	bands[0] = entity.NewBand(1, "A", 1900)
	bands[1] = entity.NewBand(2, "B", 1950)
	bands[2] = entity.NewBand(3, "C", 2000)

	return &FakeBandRepository{bandArray: bands}, nil
}

func (r *FakeBandRepository) GetBandById(id int) (entity.Band, error) {
	bandIndex := slices.IndexFunc(r.bandArray, func(band entity.Band) bool {
		return band.Id == int8(id)
	})
	if bandIndex == -1 {
		return entity.Band{}, errors.New("Not found")
	}
	return r.bandArray[bandIndex], nil
}
