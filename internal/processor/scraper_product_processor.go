package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/wenealves10/poc-asynq-colly-images/internal/queues"
	"github.com/wenealves10/poc-asynq-colly-images/internal/tasks"
)

type ScraperProductProcessor struct {
	client *asynq.Client
}

func NewScraperProductProcessor(client *asynq.Client) *ScraperProductProcessor {
	return &ScraperProductProcessor{
		client: client,
	}
}

func (p *ScraperProductProcessor) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload tasks.ScraperProductPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// Simulate a long-running task
	time.Sleep(10 * time.Second)
	fmt.Printf("Processing ScraperProduct task with URL: %s\n", payload.ProductURL)

	//todo: // send image processor task
	taskImageProcessor, err := tasks.NewImageProcessorTask("123", "https://example.com/image.jpg")
	if err != nil {
		return fmt.Errorf("failed to create image processor task: %v", err)
	}

	// Enqueue the image processor task
	if _, err := p.client.Enqueue(taskImageProcessor, asynq.Queue(queues.TypeImageProcessorQueue)); err != nil {
		return fmt.Errorf("failed to enqueue image processor task: %v", err)
	}

	return nil
}
