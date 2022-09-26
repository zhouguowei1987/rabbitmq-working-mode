package main

import (
	"fmt"
	"rabbitmq-working-mode/rbtmqcs/rbtmq_simple"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rbtmq_simple.NewRabbitMQSimple("queuetwo")
	for i := 0; i <= 1000; i++ {
		rabbitmq.PublishSimple("hello xiaobai" + strconv.Itoa(i))
		time.Sleep(1 * time.Second / 10)
		fmt.Println(i)
	}
}
