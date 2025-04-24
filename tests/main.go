package main

import (
	"fmt"

	"github.com/wenealves10/poc-asynq-colly-images/internal/scraper"
)

func main() {
	scraperMercadoLivre := scraper.NewMercadoLivreScraper("https://www.mercadolivre.com.br")
	product, err := scraperMercadoLivre.GetProductURL("https://www.mercadolivre.com.br/perfume-feminino-la-vie-est-belle-lancme-edp-100ml/p/MLB6140723")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product details: %+v\n", product)
}
