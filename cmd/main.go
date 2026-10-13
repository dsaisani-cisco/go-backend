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
			fx.Annotate(clock.New, fx.As(new(service.Clock))),
			service.New,
			api.New,
			api.NewServer,
		),
		fx.Invoke(api.Run),
	).Run()
}
