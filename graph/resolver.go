package graph

import (
	"github.com/tomoropy/fishing-with-api/usecase/presenter"
	"github.com/tomoropy/fishing-with-api/usecase/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	QS service.QueryService
	MS service.MutationService
	P  presenter.Presenter
}

func NewResolver(qs service.QueryService, ms service.MutationService, p presenter.Presenter) *Resolver {
	return &Resolver{
		QS: qs,
		MS: ms,
		P:  p,
	}
}
