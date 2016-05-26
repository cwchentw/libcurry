package libcurry

import (
	"errors"
	"github.com/pelletier/go-toml"
	"os"
	"path"
	"runtime"
	"time"
)

func IsInit() bool {
	rcPath := getRcPath()
	_, err := os.Stat(rcPath)

	if err != nil {
		return true
	} else {
		return false
	}
}

func IsDataOld() (bool, error) {
	ratesPath, err := GetCurrencyRatesPath()
	if err != nil {
		return false, err
	}

	info, err := os.Stat(ratesPath)
	if err != nil {
		return false, err
	}


	modTime := info.ModTime()
	delta := time.Now().Sub(modTime)
	if delta.Hours() > 24 {
		return true, nil
	} else {
		return false, nil
	}
}

func GetCurrencyRatesPath() (string, error) {
	dataDir, err := getDataDir()
	if err != nil {
		return "", err
	}

	currencyRatesPath := path.Join(dataDir, lastestRatesName)
	return currencyRatesPath, nil
}

func GetCurrenciesPath() (string, error) {
	dataDir, err := getDataDir()
	if err != nil {
		return "", err
	}

	currenciesPath := path.Join(dataDir, currenciesName)
	return currenciesPath, nil
}

func getDataDir() (string, error) {
	config, err := loadConfig()
	if err != nil {
		return "", err
	}

	dataDir := config["data_dir"].(string)
	if dataDir == "" {
		return "", errors.New("Unknown data directory")
	}

	_, err = os.Stat(dataDir)
	if err != nil {
		return "", err
	}

	return dataDir, nil
}

func loadConfig() (map[string]interface{}, error) {
	rcPath := getRcPath()
	tree, err := toml.LoadFile(rcPath)
	if err != nil {
		return nil, err
	}

	return tree.ToMap(), nil
}

func GetConfigPath() string {
	return getRcPath()
}

func getRcPath() string {
	return path.Join(getHomeDir(), ".currencyrc")
}

func getHomeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("APPDATA")
	} else {
		return os.Getenv("HOME")
	}
}
