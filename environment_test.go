package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMandatoryEnv(t *testing.T) {
	err := os.Setenv("test_env", "value")
	assert.Nil(t, err)

	assert.NotPanics(t, func() { GetMandatoryEnv("test_env") })
	assert.Equal(t, "value", GetMandatoryEnv("test_env"))

	err = os.Unsetenv("test_env")
	assert.Nil(t, err)
}

func TestGetMandatoryEnvPanic(t *testing.T) {
	assert.Panics(t, func() { GetMandatoryEnv("test_env") })

	err := os.Unsetenv("test_env")
	assert.Nil(t, err)
}

func TestGetDefaultEnv(t *testing.T) {
	err := os.Setenv("test_env", "value")
	assert.Nil(t, err)

	assert.Equal(t, "value", GetDefaultEnv("test_env", "test"))

	err = os.Unsetenv("test_env")
	assert.Nil(t, err)
}

func TestGetDefaultEnvPanic(t *testing.T) {
	assert.Equal(t, "test", GetDefaultEnv("test_env", "test"))

	err := os.Unsetenv("test_env")
	assert.Nil(t, err)
}
