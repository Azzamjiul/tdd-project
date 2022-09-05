package main

import (
	s "tdd/stocks"
	"testing"
)

func assertEqual(t *testing.T, expected s.Money, actual s.Money) {
	if expected != actual {
		t.Errorf("Expected [%+v], got: [%+v]", expected, actual)
	}
}

func TestMultiplication(t *testing.T) {
	fiver := s.NewMoney(5, "EUR")
	actualResult := fiver.Times(2)
	expectedResult := s.NewMoney(10, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := s.NewMoney(10, "EUR")
	actualResult := tenEuros.Times(2)
	expectedResult := s.NewMoney(20, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := s.NewMoney(4002, "KRW")
	actualResult := originalMoney.Divide(4)
	expectedResult := s.NewMoney(1000.5, "KRW")
	assertEqual(t, expectedResult, actualResult)
}

func TestAddition(t *testing.T) {
	var portofolio s.Portofolio
	var portofolioInDollars s.Money

	fiveDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "USD")
	fifteenDollars := s.NewMoney(15, "USD")

	portofolio = portofolio.Add(fiveDollars)
	portofolio = portofolio.Add(tenDollars)
	portofolioInDollars = portofolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portofolioInDollars)
}
