package main

import (
	"danielsaisani/cmd/api"
	"danielsaisani/cmd/service"
	"danielsaisani/pkg/clock"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			clock.New,
			service.New,
			api.New,
		),
	)
}
