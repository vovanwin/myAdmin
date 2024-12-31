package cmd

import (
	"go.uber.org/fx"
	"myAdmin/internal/platform/screens/users"
	"myAdmin/resources/templates"
	"myAdmin/router"
)

func Inject() fx.Option {
	return fx.Options(
		//fx.NopLogger,
		fx.Provide(
			templates.ProvideTemplate,
			router.NewRouter,
		),
		fx.Invoke(
			NewServerAdmin,
		),
		users.Module,
	)
}
