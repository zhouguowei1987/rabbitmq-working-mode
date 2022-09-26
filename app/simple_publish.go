package main

import (
	"fmt"
	"rabbitmq-working-mode/rbtmqcs/rbtmq_simple"
)

func main() {
	rabbitmq := rbtmq_simple.NewRabbitMQSimple("queueone")
	rabbitmq.PublishSimple("hello huxiaobai12345!!!")
	fmt.Println("发送成功")
}
