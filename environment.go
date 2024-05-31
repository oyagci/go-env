package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func GetMandatoryEnv(envName string) string {
	env, present := os.LookupEnv(envName)

	if !present {
		panic(fmt.Sprintf("Env var %s is mandatory", envName))
	}

	return env
}

func GetDefaultEnv(envName, defaultValue string) string {
	env, present := os.LookupEnv(envName)

	if !present {
		return defaultValue
	}

	return env
}

func GetMandatoryIntFromEnv(envName string) int {
	value := GetMandatoryEnv(envName)

	intVal, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intVal
}

func GetDefaultIntFromEnv(envName, defaultValue string) int {
	value := GetDefaultEnv(envName, defaultValue)

	intVal, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intVal
}

func GetMandatoryDurationFromEnv(value string) time.Duration {
	v := GetMandatoryEnv(value)
	d, err := time.ParseDuration(v)
	if err != nil {
		panic(err)
	}
	return d
}

func GetDefaultDurationFromEnv(duration, defaultDuration string) time.Duration {
	v := GetDefaultEnv(duration, defaultDuration)
	d, err := time.ParseDuration(v)
	if err != nil {
		panic(err)
	}
	return d
}

func GetMandatoryBoolFromEnv(name string) bool {
	value := GetMandatoryEnv(name)

	return mustParseBool(value)
}

func GetDefaultBoolFromEnv(name string, defaultValue bool) bool {
	value := GetDefaultEnv(name, strconv.FormatBool(defaultValue))

	return mustParseBool(value)
}

func mustParseBool(value string) bool {
	b, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return b
}

func GetDefaultFloat32FromEnv(envName string, defaultValue string) float32 {
	value := GetDefaultEnv(envName, defaultValue)

	floatVal, err := strconv.ParseFloat(value, 32)
	if err != nil {
		panic(err)
	}

	return float32(floatVal)
}

func GetMandatoryFloat32FromEnv(envName string) float32 {
	value := GetMandatoryEnv(envName)

	floatVal, err := strconv.ParseFloat(value, 32)
	if err != nil {
		panic(err)
	}

	return float32(floatVal)
}

func GetDefaultFloat64FromEnv(envName string, defaultValue string) float64 {
	value := GetDefaultEnv(envName, defaultValue)

	floatVal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return floatVal
}

func GetMandatoryFloat64FromEnv(envName string) float64 {
	value := GetMandatoryEnv(envName)

	floatVal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return floatVal
}
