package api

import "danielsaisani/cmd/service"

type API struct {
	timeService *service.TimeService
}

func New(service *service.TimeService) *API {
	return &API{
		timeService: service,
	}
}

// TODO: Make endpoint that returns the current time.
