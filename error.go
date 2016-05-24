package libcurry

import "errors"

func UnknownCurrency() error {
	return errors.New("Unknown currency in OpenExchangeRates")
}

func NoBaseCurrency() error {
	return errors.New("No base currency")
}

func WrongArguments() error {
	return errors.New("Wrong arguments")
}
