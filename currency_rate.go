package libcurry

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"
)

func ReadFromCurrencyRate() (float64, error) {
	rate, err := getFromToCurrencyRate("from")
	if err != nil {
		return math.NaN(), err
	}

	return rate, nil
}

func ReadToCurrencyRate() (float64, error) {
	rate, err := getFromToCurrencyRate("to")
	if err != nil {
		return math.NaN(), err
	}

	return rate, nil
}

func getFromToCurrencyRate(fromTo string) (float64, error) {
	cur, err := getFromToCurrency(fromTo)
	if err != nil {
		return math.NaN(), err
	}

	rates, err := ReadCurrencyRates()
	if err != nil {
		return math.NaN(), err
	}

	rate, ok := rates[cur]
	if ok != true {
		return math.NaN(), UnknownCurrency()
	}

	return rate, nil
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
