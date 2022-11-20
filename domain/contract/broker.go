package contract

import (
	"context"
	"github.com/jaimera/poc-services/domain/dto"
)

type PortBroker interface {
	Produce(ctx context.Context, port dto.PortDto) error
	Consume(ctx context.Context, portService PortService) error
}
