package libcurry

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

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

func getFromToCurrency(fromTo string) (string, error) {
	configTree, err := LoadConfigFile()
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

func setFromToCurrency(fromTo string, currency string) error {
	configTree, err := LoadConfigFile()
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

	err = WriteConfigFile(configTree)
	if err != nil {
		return err
	}

	return nil
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
