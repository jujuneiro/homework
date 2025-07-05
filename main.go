package main

import "fmt"

const (
	USDToEUR = 0.92
	USDToRub = 90
)

func main() {
	var EURToRUB = USDToRub / USDToEUR
	fmt.Println(EURToRUB)
}
