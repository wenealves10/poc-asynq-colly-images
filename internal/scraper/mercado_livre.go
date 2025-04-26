package scraper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"
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

	product := &model.Product{
		ID:           uuid.New().String(),
		ExternalLink: productURL,
		Images:       []string{},
	}

	var err error

	c.OnHTML("html", func(e *colly.HTMLElement) {
		product.Name = e.ChildText(".ui-pdp-title")

		// Get price fraction
		priceFraction := e.ChildText("#price > div > div.ui-pdp-price__main-container > div.ui-pdp-price__second-line.ui-pdp-price__second-line--double > span:nth-child(1) > span > span.andes-money-amount__fraction")

		// Get price cents
		priceCents := e.ChildText("#price > div > div.ui-pdp-price__main-container > div.ui-pdp-price__second-line.ui-pdp-price__second-line--double > span:nth-child(1) > span > span.andes-money-amount__cents.andes-money-amount__cents--superscript-36")

		// Convert price to float64
		var price float64
		if priceCents != "" {
			_, err = fmt.Sscanf(fmt.Sprintf("%s.%s", priceFraction, priceCents), "%f", &price)
		} else {
			_, err = fmt.Sscanf(priceFraction, "%f", &price)
		}

		if err == nil {
			product.CurrentPrice = price
		}

		// regex to extract the number of reviews

		reviewsText := e.ChildText(".ui-pdp-review__amount")
		regexp := regexp.MustCompile(`\d+`)
		reviewMatches := regexp.FindAllString(reviewsText, -1)
		if len(reviewMatches) > 0 {
			product.Reviews, err = strconv.Atoi(reviewMatches[0])
			if err != nil {
				product.Reviews = 0
			}
		} else {
			product.Reviews = 0
		}

		// Extract product ID from URL if available
		product.ExternalID = extractIDFromURL(productURL)

		// Get images
		e.ForEach(".ui-pdp-image", func(_ int, el *colly.HTMLElement) {
			product.Images = append(product.Images, el.Attr("src"))
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(productURL)

	return product, err
}

// Helper function to extract ID from URL
func extractIDFromURL(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) > 0 {
		// Get the last part of the URL which often contains the ID
		idPart := parts[len(parts)-1]
		// If ID contains a hyphen, get the part after the last hyphen
		if strings.Contains(idPart, "-") {
			subParts := strings.Split(idPart, "-")
			return subParts[len(subParts)-1]
		}
		return idPart
	}
	return ""
}
