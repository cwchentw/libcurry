package libcurry

import (
	"os"
)

func Clear() error {
	dataDir, err := getDataDir()
	if err != nil {
		return err
	}

	err = os.RemoveAll(dataDir)
	if err != nil {
		return err
	}

	rcPath := getRcPath()

	err = os.Remove(rcPath)
	if err != nil {
		return err
	}

	return nil
}
