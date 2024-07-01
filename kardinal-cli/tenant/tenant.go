package tenant

import (
	"github.com/google/uuid"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	"kardinal.cli/tenant/host_machine_directories"
	"os"
)

const (
	tenantUuidFilePermissions os.FileMode = 0644
)

func GetOrCreateUserTenantUUID() (uuid.UUID, error) {

	kardinalFkTenantUuidFilepath, err := host_machine_directories.GetKardinalFkTenantUuidFilepath()
	if err != nil {
		return uuid.UUID{}, stacktrace.Propagate(err, "An error occurred getting the tenant UUID filepath")
	}

	_, err = os.Stat(kardinalFkTenantUuidFilepath)
	if err != nil {
		if os.IsNotExist(err) {

			newUUID, err := uuid.NewRandom()
			if err != nil {
				return uuid.UUID{}, stacktrace.Propagate(err, "An error occurred generating a new UUID")
			}

			if err := os.WriteFile(kardinalFkTenantUuidFilepath, []byte(newUUID.String()), tenantUuidFilePermissions); err != nil {
				return uuid.UUID{}, stacktrace.Propagate(err, "An error occurred writing fk tenant UUID file '%v'", kardinalFkTenantUuidFilepath)
			}
			logrus.Debugf("Kardinal fk tenant UUID file saved to %v", kardinalFkTenantUuidFilepath)
			logrus.Infof("Creating new tenant UUID %s", newUUID)
			return newUUID, nil
		} else {
			return uuid.UUID{}, stacktrace.Propagate(err, "An error occurred getting fk tenant UUID file info")
		}
	}

	kardinalFkTenantUuidFileBytes, err := os.ReadFile(kardinalFkTenantUuidFilepath)
	if err != nil {
		return uuid.UUID{}, stacktrace.Propagate(err, "attempted to read file fk tenant UUID with path '%s' but failed", kardinalFkTenantUuidFilepath)
	}

	kardinalFkTenantUuidFileStr := string(kardinalFkTenantUuidFileBytes)

	parsedUuid, err := uuid.Parse(kardinalFkTenantUuidFileStr)
	if err != nil {
		return uuid.UUID{}, stacktrace.Propagate(err, "An error occurred parsing the UUID str '%s' to UUID", kardinalFkTenantUuidFileStr)
	}

	logrus.Infof("Using tenant UUID %s", kardinalFkTenantUuidFileStr)
	return parsedUuid, nil
}
