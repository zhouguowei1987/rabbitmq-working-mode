package main

import "rabbitmq-working-mode/rbtmqcs/rbtmq_sub"

func main() {
	rabbitmq := rbtmq_sub.NewRabbitMQSub("newProduct")
	rabbitmq.ConsumeSub()
}
