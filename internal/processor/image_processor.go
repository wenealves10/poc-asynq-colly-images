package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/wenealves10/poc-asynq-colly-images/internal/tasks"
)

type ImageProcessor struct {
	client *asynq.Client
}

func NewImageProcessor(client *asynq.Client) *ImageProcessor {
	return &ImageProcessor{
		client: client,
	}
}

func (p *ImageProcessor) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload tasks.ImageProcessorPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	// Simulate a long-running task
	time.Sleep(10 * time.Second)
	fmt.Printf("Processing ImageProcessor task with ProductID: %s and ImageURL: %s\n", payload.ProductID, payload.ImageURL)

	//send image uploader task
	taskImageUploader, err := tasks.NewImageUploaderTask("123", "image.jpg", "/path/to/image.jpg")
	if err != nil {
		return fmt.Errorf("failed to create image uploader task: %v", err)
	}
	// Enqueue the image uploader task
	if _, err := p.client.Enqueue(taskImageUploader); err != nil {
		return fmt.Errorf("failed to enqueue image uploader task: %v", err)
	}

	return nil
}
