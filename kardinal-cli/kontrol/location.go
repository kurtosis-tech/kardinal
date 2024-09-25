package kontrol

import (
	"fmt"
	"os"

	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"kardinal.cli/host_machine_directories"
)

const (
	KontrolLocationLocal                       = "local-kardinal-kontrol"
	KontrolLocationKloud                       = "kloud-kardinal-kontrol"
	OldKontrolLocationLocal                    = "local-minikube"
	OldKontrolLocationKloud                    = "kloud-kontrol"
	kontrolLocationFilePermissions os.FileMode = 0644
)

func SaveKontrolLocation(kontrolLocation string) error {
	kontrolLocationFilepath, err := host_machine_directories.GetKontrolLocationFilepath()
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred getting the Kontrol location filepath")
	}

	if err := os.WriteFile(kontrolLocationFilepath, []byte(kontrolLocation), kontrolLocationFilePermissions); err != nil {
		return stacktrace.Propagate(err, "An error occurred writing Kontrol location file '%v'", kontrolLocationFilepath)
	}

	return nil
}

func GetKontrolLocation() (string, error) {

	kontrolLocationFilepath, err := host_machine_directories.GetKontrolLocationFilepath()
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol location filepath")
	}

	fmt.Println(kontrolLocationFilepath)
	_, err = os.Stat(kontrolLocationFilepath)
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol location file info")
	}

	kontrolLocationFileBytes, err := os.ReadFile(kontrolLocationFilepath)
	if err != nil {
		return "", stacktrace.Propagate(err, "attempted to read file Kontrol location path '%s' but failed", kontrolLocationFilepath)
	}

	kontrolLocationFileStr := string(kontrolLocationFileBytes)
	switch kontrolLocationFileStr {
	case OldKontrolLocationLocal:
		kontrolLocationFileStr = KontrolLocationLocal
	case OldKontrolLocationKloud:
		kontrolLocationFileStr = KontrolLocationKloud
	}

	logrus.Infof("Using Kontrol location %s", kontrolLocationFileStr)
	return kontrolLocationFileStr, nil
}
