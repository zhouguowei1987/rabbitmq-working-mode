package main

import "rabbitmq-working-mode/rbtmqcs/rbtmq_routing"

func main() {
	rabbitmqTwo := rbtmq_routing.NewRabbitMQRouting("exHxb", "xiaobai_two")
	rabbitmqTwo.ConsumeRouting()
}
