package plugin

type Config struct {
	Values map[string]interface{}
}

func (c *Config) Get(key string) interface{} {
	return c.Values[key]
}

func (c *Config) Set(key string, value interface{}) {
	c.Values[key] = value
}
