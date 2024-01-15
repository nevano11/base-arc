package service

import (
	"awesomeProject/internal/core-module/repository"
	"awesomeProject/internal/entity"
)

type BandService struct {
	repository repository.BandRepository
}

func (s BandService) GetBandById(id int) (entity.Band, error) {
	return s.repository.GetBandById(id)
}

func NewBandService(bandRepository repository.BandRepository) (*BandService, error) {
	return &BandService{
		repository: bandRepository,
	}, nil
}
