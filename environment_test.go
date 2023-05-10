package env_test

import (
	"os"
	"testing"
	"time"

	"github.com/owlint/go-env"
	"github.com/stretchr/testify/assert"
)

func TestGetMandatoryEnv(t *testing.T) {
	err := os.Setenv("test_env", "value")
	assert.Nil(t, err)

	assert.NotPanics(t, func() { env.GetMandatoryEnv("test_env") })
	assert.Equal(t, "value", env.GetMandatoryEnv("test_env"))

	err = os.Unsetenv("test_env")
	assert.Nil(t, err)
}

func TestGetMandatoryEnvPanic(t *testing.T) {
	assert.Panics(t, func() { env.GetMandatoryEnv("test_env") })

	err := os.Unsetenv("test_env")
	assert.Nil(t, err)
}

func TestGetDefaultEnv(t *testing.T) {
	err := os.Setenv("test_env", "value")
	assert.Nil(t, err)

	assert.Equal(t, "value", env.GetDefaultEnv("test_env", "test"))

	err = os.Unsetenv("test_env")
	assert.Nil(t, err)
}

func TestGetDefaultEnvPanic(t *testing.T) {
	assert.Equal(t, "test", env.GetDefaultEnv("test_env", "test"))

	err := os.Unsetenv("test_env")
	assert.Nil(t, err)
}

func TestGetDefaultDurationFromEnv(t *testing.T) {
	defer os.Unsetenv("test_env")

	t.Run("env not set", func(t *testing.T) {
		assert.Equal(t, 2*time.Nanosecond, env.GetDefaultDurationFromEnv("test_env", "2ns"))
		assert.Equal(t, 2*time.Millisecond, env.GetDefaultDurationFromEnv("test_env", "2ms"))
		assert.Equal(t, 2*time.Second, env.GetDefaultDurationFromEnv("test_env", "2s"))
		assert.Equal(t, 2*time.Minute, env.GetDefaultDurationFromEnv("test_env", "2m"))
		assert.Equal(t, 2*time.Hour, env.GetDefaultDurationFromEnv("test_env", "2h"))
	})

	t.Run("invalid duration", func(t *testing.T) {
		t.Run("from default", func(t *testing.T) {
			assert.Panics(t, func() { env.GetDefaultDurationFromEnv("test_env", "invalid") })
		})

		t.Run("from env", func(t *testing.T) {
			assert.NoError(t, os.Setenv("test_env", "invalid"))
			assert.Panics(t, func() { env.GetDefaultDurationFromEnv("test_env", "2ms") })
		})
	})

	t.Run("env set", func(t *testing.T) {
		assert.NoError(t, os.Setenv("test_env", "10ns"))
		assert.Equal(t, 10*time.Nanosecond, env.GetDefaultDurationFromEnv("test_env", "2ns"))

		assert.NoError(t, os.Setenv("test_env", "10ms"))
		assert.Equal(t, 10*time.Millisecond, env.GetDefaultDurationFromEnv("test_env", "2ms"))

		assert.NoError(t, os.Setenv("test_env", "10s"))
		assert.Equal(t, 10*time.Second, env.GetDefaultDurationFromEnv("test_env", "2s"))

		assert.NoError(t, os.Setenv("test_env", "10m"))
		assert.Equal(t, 10*time.Minute, env.GetDefaultDurationFromEnv("test_env", "2m"))

		assert.NoError(t, os.Setenv("test_env", "10h"))
		assert.Equal(t, 10*time.Hour, env.GetDefaultDurationFromEnv("test_env", "2h"))
	})
}

func TestGetMandatoryDurationFromEnv(t *testing.T) {
	defer os.Unsetenv("test_env")

	t.Run("env not set", func(t *testing.T) {
		assert.Panics(t, func() { env.GetMandatoryDurationFromEnv("test_env") })
	})

	t.Run("invalid duration", func(t *testing.T) {
		assert.NoError(t, os.Setenv("test_env", "invalid"))
		assert.Panics(t, func() { env.GetMandatoryDurationFromEnv("test_env") })
	})

	t.Run("env set", func(t *testing.T) {
		assert.NoError(t, os.Setenv("test_env", "10ns"))
		assert.Equal(t, 10*time.Nanosecond, env.GetMandatoryDurationFromEnv("test_env"))

		assert.NoError(t, os.Setenv("test_env", "10ms"))
		assert.Equal(t, 10*time.Millisecond, env.GetMandatoryDurationFromEnv("test_env"))

		assert.NoError(t, os.Setenv("test_env", "10s"))
		assert.Equal(t, 10*time.Second, env.GetMandatoryDurationFromEnv("test_env"))

		assert.NoError(t, os.Setenv("test_env", "10m"))
		assert.Equal(t, 10*time.Minute, env.GetMandatoryDurationFromEnv("test_env"))

		assert.NoError(t, os.Setenv("test_env", "10h"))
		assert.Equal(t, 10*time.Hour, env.GetMandatoryDurationFromEnv("test_env"))
	})
}

func TestGetDefaultBoolFromEnv(t *testing.T) {
	defer os.Unsetenv("test_env")

	t.Run("env not set", func(t *testing.T) {
		assert.True(t, env.GetDefaultBoolFromEnv("test_env", true))
		assert.False(t, env.GetDefaultBoolFromEnv("test_env", false))
	})

	t.Run("invalid value from env", func(t *testing.T) {
		defer os.Unsetenv("test_env")

		assert.NoError(t, os.Setenv("test_env", "invalid"))
		assert.Panics(t, func() { env.GetDefaultBoolFromEnv("test_env", true) })
	})

	t.Run("env set", func(t *testing.T) {
		t.Run("true with default false", func(t *testing.T) {
			trueValues := []string{
				"1", "t", "T", "TRUE", "true", "True",
			}

			for _, value := range trueValues {
				assert.NoError(t, os.Setenv("test_env", value))
				assert.True(t, env.GetDefaultBoolFromEnv("test_env", false))
			}
		})

		t.Run("false with default true", func(t *testing.T) {
			falseValues := []string{
				"0", "f", "F", "FALSE", "false", "False",
			}

			for _, value := range falseValues {
				assert.NoError(t, os.Setenv("test_env", value))
				assert.False(t, env.GetDefaultBoolFromEnv("test_env", true))
			}
		})
	})
}
func TestGetMandatoryBoolFromEnv(t *testing.T) {
	defer os.Unsetenv("test_env")

	t.Run("invalid value", func(t *testing.T) {
		defer os.Unsetenv("test_env")

		assert.NoError(t, os.Setenv("test_env", "invalid"))
		assert.Panics(t, func() { env.GetMandatoryBoolFromEnv("test_env") })
	})

	t.Run("env set", func(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			trueValues := []string{
				"1", "t", "T", "TRUE", "true", "True",
			}

			for _, value := range trueValues {
				assert.NoError(t, os.Setenv("test_env", value))
				assert.True(t, env.GetMandatoryBoolFromEnv("test_env"))
			}
		})

		t.Run("false", func(t *testing.T) {
			falseValues := []string{
				"0", "f", "F", "FALSE", "false", "False",
			}

			for _, value := range falseValues {
				assert.NoError(t, os.Setenv("test_env", value))
				assert.False(t, env.GetMandatoryBoolFromEnv("test_env"))
			}
		})
	})
}
