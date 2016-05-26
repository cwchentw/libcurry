package libcurry

import (
	"encoding/json"
	"errors"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

func SetFromCurrency(currency string) error {
	err := setFromToCurrency("from", currency)
	if err != nil {
		return err
	}

	return nil
}

func SetToCurrency(currency string) error {
	err := setFromToCurrency("to", currency)
	if err != nil {
		return err
	}

	return nil
}

func setFromToCurrency(fromTo string, currency string) error {
	configTree, err := loadConfigFile()
	if err != nil {
		return err
	}

	currencies, err := ReadCurrencies()
	if err != nil {
		return nil
	}

	_, ok := currencies[currency]
	if ok != true {
		return UnknownCurrency()
	}

	var _currency interface{}
	_currency = currency
	configTree.Set(fromTo, _currency)

	err = writeConfigFile(configTree)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromCurrency() (string, error) {
	from, err := getFromToCurrency("from")
	if err != nil {
		return "", err
	}
	return from, nil
}

func ReadToCurrency() (string, error) {
	to, err := getFromToCurrency("to")
	if err != nil {
		return "", err
	}
	return to, nil
}

func getFromToCurrency(fromTo string) (string, error) {
	configTree, err := loadConfigFile()
	if err != nil {
		return "", err
	}

	item := configTree.Get(fromTo)

	result, ok := item.(string)
	if ok != true {
		return "", errors.New("Unable to read " + fromTo + " currency")
	}

	return result, nil
}

func ReadCurrencyRates() (map[string]float64, error) {
	rates := make(map[string]float64)

	ratesPath, err := getCurrencyRatesPath()
	if err != nil {
		return rates, err
	}

	info, err := os.Stat(ratesPath)
	if err != nil {
		return rates, err
	}

	// Update data if modification time > 24 hours
	modTime := info.ModTime()
	delta := time.Now().Sub(modTime)
	if delta.Hours() > 24 {
		err := UpdateData()
		if err != nil {
			//return rates, err
			log.Println("Data older than 24 hours")
		}
	}

	content, err := ioutil.ReadFile(ratesPath)
	if err != nil {
		return rates, err
	}

	var j interface{}
	err = json.Unmarshal(content, &j)
	if err != nil {
		return rates, err
	}

	m := j.(map[string]interface{})
	for k, v := range m {
		if k == "rates" {
			m1 := v.(map[string]interface{})
			for k1, v1 := range m1 {
				switch u := v1.(type) {
				case float64:
					rates[k1] = u
				}
			}
		}
	}

	return rates, nil
}

func ReadCurrencies() (map[string]string, error) {
	currencies := make(map[string]string)

	currenciesPath, err := getCurrenciesPath()
	if err != nil {
		return currencies, err
	}

	content, err := ioutil.ReadFile(currenciesPath)
	if err != nil {
		return currencies, err
	}

	var j interface{}
	err = json.Unmarshal(content, &j)
	if err != nil {
		return currencies, err
	}

	m := j.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			currencies[k] = vv
		default:
			// pass
		}
	}

	return currencies, nil
}

func IsInit() bool {
	rcPath := getRcPath()
	_, err := os.Stat(rcPath)

	if err != nil {
		return true
	} else {
		return false
	}
}

func loadConfigFile() (*toml.TomlTree, error) {
	var tree *toml.TomlTree
	rcPath := getRcPath()
	isInit := IsInit()
	if isInit {
		tree = toml.TreeFromMap(make(map[string]interface{}))
	} else {
		_tree, err := toml.LoadFile(rcPath)

		if err != nil {
			return _tree, err
		}

		tree = _tree
	}
	return tree, nil
}

func writeConfigFile(tree *toml.TomlTree) error {
	rcPath := getRcPath()
	f, err := os.Create(rcPath)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(tree.ToString())

	return nil
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

func getRcPath() string {
	return path.Join(getHomeDir(), ".currencyrc")
}

func getHomeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("HOMEPATH")
	} else {
		return os.Getenv("HOME")
	}
}
