package libcurry

import (
	"encoding/json"
	"io/ioutil"
	"math"
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

	ratesPath, err := GetCurrencyRatesPath()
	if err != nil {
		return rates, err
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
