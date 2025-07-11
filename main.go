package main

import "fmt"

const (
	USDToEUR = 0.92
	USDToRub = 90
)

func main() {
	var EURToRUB = USDToRub / USDToEUR
	fmt.Println("Курс евро к рублю: ", EURToRUB)

	summa, ishod, celevaya := vvod()
	result := raschet(summa, ishod, celevaya)
	fmt.Println(result)
}

func vvod() (float64, string, string) {
	var ishod string
	var celevaya string
	var summa float64
	fmt.Println("Введите сумму: ")
	fmt.Scan(&summa)
	fmt.Println("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&ishod)
	fmt.Println("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&celevaya)
	return summa, ishod, celevaya
}

func raschet(summa float64, ishod, celevaya string) float64 {
	var result float64
	switch ishod {
	case "USD":
		if celevaya == "EUR" {
			result = summa * USDToEUR
		} else if celevaya == "RUB" {
			result = summa * USDToRub
		}
	case "EUR":
		if celevaya == "EUR" {
			result = summa / USDToEUR
		} else if celevaya == "RUB" {
			result = summa * USDToRub / USDToEUR
		}
	case "RUB":
		if celevaya == "USD" {
			result = summa / USDToEUR
		} else if celevaya == "RUB" {
			result = summa * USDToEUR / USDToRub
		}
	}
	return result
}
