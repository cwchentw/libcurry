package libcurry

import (
	"errors"
	"github.com/pelletier/go-toml"
	"os"
	"path"
	"runtime"
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

func getCurrencyRatesPath() (string, error) {
	dataDir, err := getDataDir()
	if err != nil {
		return "", err
	}

	currencyRatesPath := path.Join(dataDir, lastestRatesName)
	return currencyRatesPath, nil
}

func getCurrenciesPath() (string, error) {
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
