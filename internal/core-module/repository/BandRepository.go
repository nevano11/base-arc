package repository

import "awesomeProject/internal/entity"

type BandRepository interface {
	GetBandById(id int) (entity.Band, error)
}
