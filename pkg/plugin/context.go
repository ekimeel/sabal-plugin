package plugin

import "database/sql"

type Context struct {
	PluginManager *PluginManager
	DB            *sql.DB
	Services      map[string]interface{}
}

func (c *Context) GetDB() (*sql.DB, bool) {
	if c.DB == nil {
		return nil, false
	}
	return c.DB, true
}

func (c *Context) GetService(name string) (interface{}, bool) {
	v := c.Services[name]
	if v == nil {
		return nil, false
	}
	return v, true
}
