package main

import (
	"context"
	"fmt"
	"log"
	"redis-demo/model"
	"redis-demo/redis"
	"time"
)

func main() {
	stringFuncs()

}

func nxFuncs() {
	// 链接redis
	cli, err := redis.GetConnect()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	setnx, err := cli.SetNX(ctx, "key1", "value1", 10*time.Second).Result()
	if setnx {
		fmt.Println("setnx success")

	} else {

	}
}

func stringFuncs() {
	// 链接redis
	cli, err := redis.GetConnect()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// set1
	res := cli.Set(ctx, "key1", "value2", 0)
	fmt.Println(res.String())
	// get1
	get1, err := cli.Get(ctx, "key1").Result()
	fmt.Println(get1)

	// set2
	res = cli.Set(ctx, "key2", &model.User{
		Id:       1,
		Name:     "test",
		Password: "123456",
	}, 0)
	fmt.Println(res.String())
	// get2
	get2, err := cli.Get(ctx, "key2").Result()
	fmt.Println(get2)

	// 过期
	//res = cli.Set(ctx, "key3", "value3", 5*time.Second)
	//tk := time.NewTicker(time.Second)
	//for {
	//	select {
	//	case <-tk.C:
	//		fmt.Println(res.String())
	//		get3, err := cli.Get(ctx, "key3").Result()
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		fmt.Println(get3)
	//	}
	//}

	result, err := cli.Do(ctx, "SET", "key4", "value4", "EX", 5).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("key4", result.(string))

	// 订阅channel1这个channel
	sub := cli.Subscribe(ctx, "channel1")
	// sub.Channel() 返回go channel，可以循环读取redis服务器发过来的消息
	for msg := range sub.Channel() {
		// 打印收到的消息
		fmt.Println(msg.Channel)
		fmt.Println(msg.Payload)
	}

}
