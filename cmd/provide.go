package cmd

import (
	"go.uber.org/fx"
	"myAdmin/internal/platform/screens/devices"
	"myAdmin/resources/templates"
	"myAdmin/router"
)

func Inject() fx.Option {
	return fx.Options(
		//fx.NopLogger,
		fx.Provide(
			templates.TemplateRenderer,
			NewServerAdmin,
		),

		//users.Module,
		devices.Module,

		fx.Invoke(
			router.NewRouter,
			ServerAdminRun,
		),
	)
}
