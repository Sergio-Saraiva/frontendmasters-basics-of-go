package main

import (
	"fmt"
	"sync"

	"frontendmasters.com/courses/go-basics/cryptomasters/api"
)

func main() {

	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(currency string) {
			getCurrencyData(currency)
			wg.Done()
		}(currency)
	}
	wg.Wait()

}

func getCurrencyData(currency string) {
	res, err := api.GetRate(currency)
	if err != nil {
		panic(err)

	}

	fmt.Printf("The rate for %v is %.2f\n", res.Currency, res.Price)
}
