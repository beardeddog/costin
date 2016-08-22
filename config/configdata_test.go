package config_test

import (
	"costin/config"
	"testing"
)

var (
	values = []string{"B", "C"}
	cfg    = config.Config{
		Type:       "config",
		Collection: "stamps",
		Field:      "Gum",
		Values:     values,
	}
)

func TestConfigString(t *testing.T) {
	t.Log(cfg.String())
}
func TestConfigJSON(t *testing.T) {
	t.Log(cfg.JSON())
}
