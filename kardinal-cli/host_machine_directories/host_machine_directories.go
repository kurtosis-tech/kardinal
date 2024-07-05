package host_machine_directories

import (
	"github.com/adrg/xdg"
	"github.com/kurtosis-tech/stacktrace"
	"path"
)

const (
	applicationDirname = "kardinal"
	fkTenantUUID       = "fk-tenant-uuid"
	kontrolLocation    = "kontrol-location"
)

func GetKardinalFkTenantUuidFilepath() (string, error) {
	xdgRelFilepath := getRelativeFilepathForXDG(fkTenantUUID)
	fkTenantUuidFilepath, err := xdg.DataFile(xdgRelFilepath)
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the fk tenant UUID filepath from relative path '%v'", xdgRelFilepath)
	}
	return fkTenantUuidFilepath, nil
}

func GetKontrolLocation() (string, error) {
	xdgRelFilepath := getRelativeFilepathForXDG(kontrolLocation)
	kontrolLocationFilepath, err := xdg.DataFile(xdgRelFilepath)
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred getting the Kontrol location filepath from relative path '%v'", xdgRelFilepath)
	}
	return kontrolLocationFilepath, nil
}

// Joins the "kardinal" app directory in front of whichever filepath
func getRelativeFilepathForXDG(filepathRelativeToKurtosisDir string) string {
	return path.Join(applicationDirname, filepathRelativeToKurtosisDir)
}
