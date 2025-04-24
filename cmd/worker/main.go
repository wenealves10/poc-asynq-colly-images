package main

import (
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/wenealves10/poc-asynq-colly-images/internal/configs"
	"github.com/wenealves10/poc-asynq-colly-images/internal/processor"
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
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
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
