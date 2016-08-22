package config

import (
	"encoding/json"
	"fmt"
)

// Config ...
type Config struct {
	Type       string   `json:"type"`
	Collection string   `json:"collection"`
	Field      string   `json:"field"`
	Values     []string `json:"values"`
}

// String ...
func (c *Config) String() string {
	return fmt.Sprintf("%s, %s, %s, %v", c.Type, c.Collection, c.Field, c.Values)
}

// JSON ...
func (c *Config) JSON() string {
	js, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("error MarshalIndent the Config struct")
		return ""
	}
	return string(js)
}
