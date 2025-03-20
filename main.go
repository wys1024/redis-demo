package main

import (
	"context"
	"fmt"
	"log"
	"redis-demo/model"
	"redis-demo/redis"
)

func main() {
	//链接redis
	cli, err := redis.GetConnect()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	// set1
	res := cli.Set(ctx, "key1", "value2", 0)
	fmt.Println(res.String())
	fmt.Println(cli.Get(ctx, "key1").String())

	//set2
	res = cli.Set(ctx, "key2", &model.User{
		Id:       1,
		Name:     "test",
		Password: "123456",
	}, 0)
	fmt.Println(res.String())
	fmt.Println(cli.Get(ctx, "key2").String())

}
