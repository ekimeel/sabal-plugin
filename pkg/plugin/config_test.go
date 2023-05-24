package plugin

import (
	"testing"
)

func TestConfig_Set(t *testing.T) {
	c := &Config{
		Values: make(map[string]interface{}),
	}

	// set a string value
	c.Set("key1", "value1")

	if c.Values["key1"] != "value1" {
		t.Errorf("expected 'value1', got %v", c.Values["key1"])
	}

	// set an integer value
	c.Set("key2", 1234)

	if c.Values["key2"] != 1234 {
		t.Errorf("expected 1234, got %v", c.Values["key2"])
	}

	// overwrite an existing value
	c.Set("key1", "newValue1")

	if c.Values["key1"] != "newValue1" {
		t.Errorf("expected 'newValue1', got %v", c.Values["key1"])
	}
}

func TestConfig_Get(t *testing.T) {
	c := &Config{
		Values: map[string]interface{}{
			"key1": "value1",
			"key2": 1234,
		},
	}

	v := c.Get("key1")

	if v != "value1" {
		t.Errorf("expected 'value1', got %v", v)
	}

	v = c.Get("key2")

	if v != 1234 {
		t.Errorf("expected 1234, got %v", v)
	}

	// try to get a non-existent key
	v = c.Get("nonExistentKey")

	if v != nil {
		t.Errorf("expected nil, got %v", v)
	}
}
