package kontrol

import (
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"kardinal.cli/host_machine_directories"
	"os"
)

const (
	kontrolLocationFilePermissions os.FileMode = 0644
)

func SaveKontrolLocation(kontrolLocation string) error {
	kontrolLocationFilepath, err := host_machine_directories.GetKontrolLocation()
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred getting the Kontrol location filepath")
	}

	if err := os.WriteFile(kontrolLocationFilepath, []byte(kontrolLocation), kontrolLocationFilePermissions); err != nil {
		return stacktrace.Propagate(err, "An error occurred writing Kontrol location file '%v'", kontrolLocationFilepath)
	}

	return nil
}

func GetKontrolLocation() (string, error) {

	kontrolLocationFilepath, err := host_machine_directories.GetKontrolLocation()
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol location filepath")
	}

	_, err = os.Stat(kontrolLocationFilepath)
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol location  file info")
	}

	kontrolLocationFileBytes, err := os.ReadFile(kontrolLocationFilepath)
	if err != nil {
		return "", stacktrace.Propagate(err, "attempted to read file Kontrol location path '%s' but failed", kontrolLocationFilepath)
	}

	kontrolLocationFileStr := string(kontrolLocationFileBytes)

	logrus.Infof("Using Kontrol location %s", kontrolLocationFileStr)
	return kontrolLocationFileStr, nil
}
