package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrice        []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("An error occured!")
		fmt.Println(err)
		return
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the content failed!")
		fmt.Println(err)
		file.Close()
		return
	}
}

func (job TaxIncludedPriceJob) Proccess() {
	result := make(map[string]float64)
	for _, price := range job.InputPrice {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

func NewTaxInclududedJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrice: []float64{10, 20, 30},
		TaxRate:    taxRate,
	}
}
