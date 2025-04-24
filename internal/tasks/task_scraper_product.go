package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

type ScraperProductPayload struct {
	ProductURL string `json:"product_url"`
}

func NewScraperProductTask(productURL string) (*asynq.Task, error) {
	payload, err := json.Marshal(ScraperProductPayload{ProductURL: productURL})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeScraperProduct, payload), nil
}
