package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

type ImageProcessorPayload struct {
	ProductID string `json:"product_id"`
	ImageURL  string `json:"image_url"`
}

func NewImageProcessorTask(productID, imageURL string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageProcessorPayload{ProductID: productID, ImageURL: imageURL})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeImageProcessor, payload), nil
}
