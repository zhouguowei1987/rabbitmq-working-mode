package main

import (
	"fmt"
	"rabbitmq-working-mode/rbtmqcs/rbtmq_sub"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rbtmq_sub.NewRabbitMQSub("newProduct")
	for i := 0; i < 1000; i++ {
		rabbitmq.PublishSub("订阅模式生成第" + strconv.Itoa(i) + "条数据")
		fmt.Println("订阅模式生成第" + strconv.Itoa(i) + "条数据")
		time.Sleep(1 * time.Second / 10)
	}
}
