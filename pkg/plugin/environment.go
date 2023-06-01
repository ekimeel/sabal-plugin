package plugin

const (
	DbVar                  = "storage/database"
	MetricServiceClientVar = "service/MetricServiceClient"
	PointServiceClient     = "service/PointServiceClient"
	PluginManagerVar       = "manager/PluginManager"
)

type Environment struct {
	registry map[string]any
}

func (e *Environment) Get(name string) (any, bool) {
	if e.registry == nil {
		return nil, false
	}

	v := e.registry[name]
	if v == nil {
		return nil, false
	}
	return v, true
}

func (e *Environment) Set(name string, value any) {

	if e.registry == nil {
		e.registry = make(map[string]any, 0)
	}

	e.registry[name] = value
}
