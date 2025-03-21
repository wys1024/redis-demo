package main

import (
	"context"
	"log"
	"redis-demo/redis"
)

func main() {
	ctx := context.Background()
	cli, err := redis.GetConnect()
	if err != nil {
		log.Fatal(err)
	}
	cli.Publish(ctx, "channel1", "hello")

}
