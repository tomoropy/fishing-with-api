package registry

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/tomoropy/fishing-with-api/graph"
	"github.com/tomoropy/fishing-with-api/infra"
	"github.com/tomoropy/fishing-with-api/usecase/presenter"
	"github.com/tomoropy/fishing-with-api/usecase/service"
)

type Register interface {
	NewResolver(context.Context) (*graph.Resolver, error)
}

type register struct {
	db *sqlx.DB
}

func NewRegister(db *sqlx.DB) Register {
	return &register{
		db: db,
	}
}

// DI in Resolver
func (r *register) NewResolver(ctx context.Context) (*graph.Resolver, error) {
	ur := infra.NewUserRepository(r.db)

	qs := service.NewQueryService(ur)
	ms := service.NewMutationService(ur)
	p := presenter.NewPresenter()

	resolver := graph.NewResolver(qs, ms, p)
	return resolver, nil
}
