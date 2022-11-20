package contract

import (
	"context"
	"github.com/jaimera/poc-services/domain/dto"
	"github.com/jaimera/poc-services/domain/entity"
)

type PortService interface {
	GetByCode(ctx context.Context, code string) (*entity.Port, error)
	Insert(ctx context.Context, ports []dto.PortDto) error
}
