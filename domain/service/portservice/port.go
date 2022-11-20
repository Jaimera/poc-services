package portservice

import (
	"context"
	"github.com/jaimera/poc-services/domain/contract"
	"github.com/jaimera/poc-services/domain/dto"
	"github.com/jaimera/poc-services/domain/entity"
	"github.com/sirupsen/logrus"
	"strings"
)

type PortService struct {
	repository contract.DataManager
	log        *logrus.Entry
}

func NewPortService(
	data contract.DataManager,
	log *logrus.Entry,
) PortService {
	return PortService{
		repository: data,
		log:        log,
	}
}

// GetByCode fetch a port by it's code
func (s PortService) GetByCode(ctx context.Context, code string) (*entity.Port, error) {
	port, err := s.repository.Port().FetchByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return port, nil
}

// Insert this method receives a list of ports and attempt to upsert it into the database
func (s PortService) Insert(ctx context.Context, ports []dto.PortDto) error {
	for _, dto := range ports {
		port := entity.Port{
			Slug: dto.Slug,
			Name: dto.Name,
			Coordinates: entity.LatLng{
				Latitude:  dto.Latitude,
				Longitude: dto.Longitude,
			},
			City:     dto.City,
			Country:  dto.Country,
			Province: dto.Province,
			Timezone: dto.Timezone,
			Code:     dto.Code,
		}
		if dto.Alias != nil && len(dto.Alias) > 0 {
			alias := strings.Join(dto.Alias, ";")
			port.Alias = &alias
		}
		if dto.Alias != nil && len(dto.Regions) > 0 {
			regions := strings.Join(dto.Regions, ";")
			port.Regions = &regions
		}
		port.Unlocs = strings.Join(dto.Unlocs, ";")

		_, err := s.repository.Port().Upsert(ctx, port)
		if err != nil {
			return err
		}
	}

	return nil
}
