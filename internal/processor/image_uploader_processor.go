package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/wenealves10/poc-asynq-colly-images/internal/tasks"
)

type ImageUploaderProcessor struct {
	client *asynq.Client
}

func NewImageUploaderProcessor(client *asynq.Client) *ImageUploaderProcessor {
	return &ImageUploaderProcessor{
		client: client,
	}
}

func (p *ImageUploaderProcessor) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload tasks.ImageUploaderPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	// Simulate a long-running task
	time.Sleep(10 * time.Second)
	fmt.Printf("Processing ImageUploader task with ProductID: %s and ImageURL: %s\n", payload.ProductID, payload.Filename)
	return nil
}
