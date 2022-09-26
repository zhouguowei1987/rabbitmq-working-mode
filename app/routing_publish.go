package main

import (
	"fmt"
	"rabbitmq-working-mode/rbtmqcs/rbtmq_routing"
	"strconv"
	"time"
)

func main() {
	rabbitmqOne := rbtmq_routing.NewRabbitMQRouting("exHxb", "xiaobai_one")
	rabbitmqTwo := rbtmq_routing.NewRabbitMQRouting("exHxb", "xiaobai_two")

	for i := 0; i < 1000; i++ {
		rabbitmqOne.PublishRouting("hello xiaobai one" + strconv.Itoa(i))
		rabbitmqTwo.PublishRouting("hello xiaobai two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second / 10)
		fmt.Println(i)
	}
}
