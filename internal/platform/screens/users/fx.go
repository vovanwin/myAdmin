package users

import (
	"go.uber.org/fx"
)

var Module = fx.Module("userScreen",
	fx.Invoke(UserScreen),
)
