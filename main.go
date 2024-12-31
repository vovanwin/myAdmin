package main

import (
	"go.uber.org/fx"
	"myAdmin/cmd"
)

func main() {
	fx.New(cmd.Inject()).Run()
}
