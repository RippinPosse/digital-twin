package resolver

import (
	"api/internal/graph/dataloader"
	"api/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service    *service.Service
	dataloader *dataloader.Dataloader
}

func New(service *service.Service, dataloader *dataloader.Dataloader) *Resolver {
	return &Resolver{
		service:    service,
		dataloader: dataloader,
	}
}
