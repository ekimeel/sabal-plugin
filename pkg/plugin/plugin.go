package plugin

import (
	"context"
)

// Plugin is an interface for different kinds of plugins that can
// process events, can be installed in an environment, and have a name.
type Plugin interface {
	// Process is a method that is supposed to process an Event in a certain way.
	// The implementation of this method depends on the specific plugin.
	// The context.Context argument can be used for cancellation, timeouts, and
	// to pass other context-specific values.
	Process(context context.Context, event Event)

	// Install is a method that is supposed to install the plugin in the given environment.
	// The method receives an Environment object, which it can use to perform the installation.
	// The exact steps for installation are dependent on the plugin and the environment.
	// If the installation is successful, the method should return nil. Otherwise, it
	// should return an error that describes what went wrong.
	// The context.Context argument can be used for cancellation, timeouts, and
	// to pass other context-specific values.
	Install(context context.Context, env *Environment) error

	// Name is a method that returns the name of the plugin. The name is expected to be
	// unique among all plugins, and it can be used for logging, error reporting, etc.
	Name() string

	// Type is method that returns the type of plugin.
	Type() PluginType
}
