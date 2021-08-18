package env

import (
	"fmt"
	"os"
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
