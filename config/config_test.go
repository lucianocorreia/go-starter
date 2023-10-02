package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	_, err := NewConfig()
	require.NoError(t, err)

	env := environment("abc")
	SwitchEnvironment(env)
	cfg, err := NewConfig()
	require.NoError(t, err)
	assert.Equal(t, env, cfg.App.Environment)
}
