package main

import "rabbitmq-working-mode/rbtmqcs/rbtmq_routing"

func main() {
	rabbitmqOne := rbtmq_routing.NewRabbitMQRouting("exHxb", "xiaobai_one")
	rabbitmqOne.ConsumeRouting()
}
