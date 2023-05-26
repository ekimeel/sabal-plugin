package plugin

type Config struct {
	Values map[string]interface{}
}

// Get returns the value if present
func (c *Config) Get(key string) interface{} {
	return c.Values[key]
}

// Set sets or replaces a values
func (c *Config) Set(key string, value interface{}) {
	c.Values[key] = value
}
