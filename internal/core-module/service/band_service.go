package service

import (
	"awesomeProject/internal/entity"
)

type BandRepository interface {
	GetBandById(id int) (entity.Band, error)
}

type BandService struct {
	repository BandRepository
}

func (s BandService) GetBandById(id int) (entity.Band, error) {
	return s.repository.GetBandById(id)
}

func NewBandService(bandRepository BandRepository) (*BandService, error) {
	return &BandService{
		repository: bandRepository,
	}, nil
}
