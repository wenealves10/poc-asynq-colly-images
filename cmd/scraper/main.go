package main

import (
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/wenealves10/poc-asynq-colly-images/internal/configs"
	"github.com/wenealves10/poc-asynq-colly-images/internal/tasks"
)

func main() {
	redisAddr := fmt.Sprintf("%s:%s", configs.REDIS_HOST, configs.REDIS_PORT)
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Username: configs.REDIS_USERNAME,
		Password: configs.REDIS_PASSWORD,
	})
	defer client.Close()

	// Enqueue a task
	task, err := tasks.NewScraperProductTask("https://example.com/product")
	if err != nil {
		log.Fatalf("failed to create task: %v", err)
	}

	// Enqueue the task
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("failed to enqueue task: %v", err)
	}

	log.Printf("Enqueued task: ID=%s, Type=%s, Queue=%s", info.ID, info.Type, info.Queue)
}
