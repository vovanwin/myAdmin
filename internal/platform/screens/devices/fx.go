package devices

import (
	"go.uber.org/fx"
)

var Module = fx.Module("devicesScreen",
	fx.Invoke(RegisterRoutes),
)
