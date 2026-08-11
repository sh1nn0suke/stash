package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_GetAllPluginConfiguration(t *testing.T) {
	i := InitializeEmpty()

	assert.Equal(t, i.GetAllPluginConfiguration(), map[string]map[string]any{})

	i.SetPluginConfiguration("plugin1", map[string]any{"key1": "value1"})

	assert.Equal(t, map[string]map[string]any{
		"plugin1": {"key1": "value1"},
	}, i.GetAllPluginConfiguration())

	i.SetPluginConfiguration("plugin2", map[string]any{"key2": "value2"})

	assert.Equal(t, map[string]map[string]any{
		"plugin1": {"key1": "value1"},
		"plugin2": {"key2": "value2"},
	}, i.GetAllPluginConfiguration())

	// ensure SetPluginConfiguration overwrites existing configuration
	i.SetPluginConfiguration("plugin2", map[string]any{"key3": "value3"})

	assert.Equal(t, map[string]map[string]any{
		"plugin1": {"key1": "value1"},
		"plugin2": {"key3": "value3"},
	}, i.GetAllPluginConfiguration())
}
