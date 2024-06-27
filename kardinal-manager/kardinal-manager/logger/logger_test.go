package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestSetLoggerLevel_Success(t *testing.T) {
	validLogLevels := []string{"panic", "fatal", "error", "warning", "info", "debug", "trace"}

	for _, logLevelStr := range validLogLevels {
		// lowercase
		os.Setenv(loggerLogLeverEnvVarKey, logLevelStr)
		err := SetLevel()
		require.NoError(t, err)
		logrusLevel := logrus.GetLevel()

		require.Equal(t, logLevelStr, logrusLevel.String())

		// uppercase
		os.Setenv(loggerLogLeverEnvVarKey, strings.ToUpper(logLevelStr))
		err = SetLevel()
		require.NoError(t, err)
		logrusLevel = logrus.GetLevel()

		require.Equal(t, logLevelStr, logrusLevel.String())
	}

}

func TestSetLoggerLevel_WrongLevel(t *testing.T) {
	levelStr := "wrong"
	os.Setenv(loggerLogLeverEnvVarKey, levelStr)
	err := SetLevel()
	require.Error(t, err)
}
