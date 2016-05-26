package libcurry

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func UpdateData() error {
	configTree, err := LoadConfigFile()
	if err != nil {
		return err
	}

	dataDirObj := configTree.Get("data_dir")
	dataDir, ok := dataDirObj.(string)
	if ok != true {
		return errors.New("Failed to get data directory")
	}

	if dataDir == "" {
		return errors.New("Unknown data directory")
	}

	_, err = os.Stat(dataDir)
	if err != nil {
		os.MkdirAll(dataDir, 0755)
	}

	err = updateCurrencies(dataDir)
	if err != nil {
		return err
	}

	appIDObj := configTree.Get("app_id")
	appID, ok := appIDObj.(string)
	if ok != true {
		return errors.New("Failed to get app id")
	}

	if appID == "" {
		return errors.New("Unknown app id")
	}

	err = updateCurrencyRates(dataDir, appID)
	if err != nil {
		return err
	}

	return nil
}

func updateCurrencyRates(dir string, appID string) error {
	res, err := http.Get(baseUrl + lastestRatesName + "?app_id=" + appID)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = writeBody(&res.Body, dir, lastestRatesName)
	if err != nil {
		return err
	}

	return nil
}

func updateCurrencies(dir string) error {
	res, err := http.Get(baseUrl + currenciesName)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = writeBody(&res.Body, dir, currenciesName)
	if err != nil {
		return err
	}

	return nil
}

func writeBody(body *io.ReadCloser, dir string, file string) error {
	content, err := ioutil.ReadAll(*body)
	if err != nil {
		return err
	}

	filePath := path.Join(dir, file)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(content)

	return nil
}
