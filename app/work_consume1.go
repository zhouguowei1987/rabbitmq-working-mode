package main

import (
	"rabbitmq-working-mode/rbtmqcs/rbtmq_simple"
)

func main() {
	rabbitmq := rbtmq_simple.NewRabbitMQSimple("queuetwo")
	rabbitmq.ConsumeSimple()
}
