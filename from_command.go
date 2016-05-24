package libcurry

import (
	"fmt"
	"strconv"
)

func ParseFromCmd(args []string) error {
	if len(args) == 4 {
		if args[1] != "from" {
			return WrongArguments()
		}
		return fromToBase(args[2], args[3])
	} else if len(args) == 6 {
		if args[1] != "from" && args[3] != "to" {
			return WrongArguments()
		}
		return fromTo(args[2], args[4], args[5])
	} else {
		return WrongArguments()
	}
}

func fromToBase(cur string, amountStr string) error {
	configTree, err := loadConfigFile()
	if err != nil {
		return err
	}

	base := configTree.Get("base")
	if base == nil {
		return NoBaseCurrency()
	}

	return fromTo(cur, base.(string), amountStr)
}

func fromTo(cur1 string, cur2 string, amountStr string) error {
	rates, err := ReadCurrencyRates()

	curFrom, ok := rates[cur1]
	if ok != true {
		return UnknownCurrency()
	}

	curTo, ok := rates[cur2]
	if ok != true {
		return UnknownCurrency()
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return err
	}

	fmt.Printf("%.2f\n", amount * curTo / curFrom)

	return nil
}
