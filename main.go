package main

import (
	"exampl.com/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxInclududedJob(taxRate)
		priceJob.Proccess()
	}

}
