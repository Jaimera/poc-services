package contract

import (
	"context"
	"database/sql"
	"github.com/jaimera/poc-services/domain/entity"
)

type DataManager interface {
	RepoManager
	Begin() (*sql.Tx, error)
	Close() error
}

type RepoManager interface {
	Port() PortRepository
}

type PortRepository interface {
	FetchByCode(context context.Context, code string) (*entity.Port, error)
	Upsert(context context.Context, port entity.Port) (uint32, error)
}
