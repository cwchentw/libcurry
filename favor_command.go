package libcurry

import (
	"errors"
)

func SetFavor(cur string) error {
	currencies, err := ReadCurrencies()
	if err != nil {
		return err
	}

	_, ok := currencies[cur]
	if ok != true {
		return errors.New("Unknown currency in OpenExchangeRates")
	}

	configTree, err := loadConfigFile()
	if err != nil {
		return err
	}

	var _cur interface{} = cur
	configTree.Set("base", _cur)

	err = writeConfigFile(configTree)
	if err != nil {
		return err
	}

	return nil
}
