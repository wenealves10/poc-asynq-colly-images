package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

type ImageUploaderPayload struct {
	ProductID string `json:"product_id"`
	Filename  string `json:"filename"`
	Filepath  string `json:"filepath"`
}

func NewImageUploaderTask(productID, filename, filepath string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageUploaderPayload{ProductID: productID, Filename: filename, Filepath: filepath})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeImageUploader, payload), nil
}
