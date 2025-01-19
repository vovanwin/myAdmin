package docs

import (
	"go.uber.org/fx"
)

var Module = fx.Module("docsScreen",
	fx.Invoke(DocsScreen),
)
