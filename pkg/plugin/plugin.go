package plugin

import (
	"context"
	"plugin"
)

type Plugin interface {
	Process(context context.Context, event Event)
	Install(env *Environment) error
	Name() string
}

type PluginManager interface {
	RunChannel() chan plugin.Plugin
}
