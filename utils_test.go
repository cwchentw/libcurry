package libcurry

import (
	"log"
	"testing"
)

func TestReadFromCurrency(t *testing.T) {
	_, err := ReadFromCurrency()
	if err != nil {
		errSet := SetFromCurrency("EUR")
		if errSet != nil {
			log.Fatalln(errSet)
		}
	}

	from, err := ReadFromCurrency()
	if err != nil {
		t.Fatal(err)
	}

	if from == "" {
		t.Error("Empty from currency")
	}
}

func TestReadToCurrency(t *testing.T) {
	_, err := ReadToCurrency()
	if err != nil {
		errSet := SetToCurrency("USD")
		if errSet != nil {
			log.Fatalln(errSet)
		}
	}

	to, err := ReadToCurrency()
	if err != nil {
		t.Fatal(err)
	}

	if to == "" {
		t.Error("Empty to currency")
	}
}
