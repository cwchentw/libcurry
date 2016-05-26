package libcurry

import (
	"testing"
)

func TestReadFromCurrencyRate(t *testing.T) {
	_, err := ReadFromCurrencyRate()
	if err != nil {
		t.Fatal("Unable to read from currency rate")
	}
}

func TestReadToCurrencyRate(t *testing.T) {
	_, err := ReadToCurrencyRate()
	if err != nil {
		t.Fatal("Unable to read to currency rate")
	}
}
