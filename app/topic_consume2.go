package main

import "rabbitmq-working-mode/rbtmqcs/rbtmq_topic"

func main() {
	rabbitmqOne := rbtmq_topic.NewRabbitMQTopic("hxbExc", "huxiaobai.*.cs")
	rabbitmqOne.ConsumeTopic()
}
