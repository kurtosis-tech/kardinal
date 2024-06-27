package utils

import (
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func GetFromEnvVar(
	key string,
	subject string,
) (string, error) {
	value := os.Getenv(key)
	if len(value) < 1 {
		return "", stacktrace.NewError("No '%s' env var was found. Must be provided as env var %s", subject, key)
	}
	logrus.Debugf("Successfully loaded env var '%s'", subject)
	return value, nil
}

func GetIntFromEnvVar(
	key string,
	subject string,
) (int, error) {
	strVal, err := GetFromEnvVar(key, subject)
	if err != nil {
		return 0, stacktrace.Propagate(err, "An error occurred getting env var with key '%s'", key)
	}
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		return 0, stacktrace.Propagate(err, "An error occurred converting string value '%s' to int", strVal)
	}
	return intVal, nil
}
