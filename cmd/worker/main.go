package main

import (
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/wenealves10/poc-asynq-colly-images/internal/configs"
	"github.com/wenealves10/poc-asynq-colly-images/internal/processor"
	"github.com/wenealves10/poc-asynq-colly-images/internal/queues"
	"github.com/wenealves10/poc-asynq-colly-images/internal/tasks"
)

func main() {

	redisAddr := fmt.Sprintf("%s:%s", configs.REDIS_HOST, configs.REDIS_PORT)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     redisAddr,
			Username: configs.REDIS_USERNAME,
			Password: configs.REDIS_PASSWORD,
		},
		asynq.Config{
			Concurrency: queues.Concurrency,
			Queues: map[string]int{
				queues.TypeScraperProductQueue: queues.ConcurrencyScraperProduct,
				queues.TypeImageProcessorQueue: queues.ConcurrencyImageProcessor,
				queues.TypeImageUploaderQueue:  queues.ConcurrencyImageUploader,
			},
		},
	)

	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Username: configs.REDIS_USERNAME,
		Password: configs.REDIS_PASSWORD,
	})
	defer client.Close()

	mux := asynq.NewServeMux()
	mux.Handle(tasks.TypeScraperProduct, processor.NewScraperProductProcessor(client))
	mux.Handle(tasks.TypeImageProcessor, processor.NewImageProcessor(client))
	mux.Handle(tasks.TypeImageUploader, processor.NewImageUploaderProcessor(client))

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
