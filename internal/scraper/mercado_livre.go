package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/wenealves10/poc-asynq-colly-images/internal/model"
)

type scraperMercadoLivre struct {
	domain string
}

func NewMercadoLivreScraper(domain string) *scraperMercadoLivre {
	return &scraperMercadoLivre{
		domain: domain,
	}
}

func (s *scraperMercadoLivre) GetProductURL(productURL string) (*model.Product, error) {
	c := colly.NewCollector()

	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println("Title:", e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(productURL)

	return nil, nil
}
