package main

import (
	"fmt"
	"rabbitmq-working-mode/rbtmqcs/rbtmq_topic"
	"strconv"
	"time"
)

func main() {
	rabbitmqOne := rbtmq_topic.NewRabbitMQTopic("hxbExc", "huxiaobai.one")
	rabbitmqTwo := rbtmq_topic.NewRabbitMQTopic("hxbExc", "huxiaobai.two.cs")

	for i := 0; i < 1000; i++ {
		rabbitmqOne.PublishTopic("hello huxiaobai one" + strconv.Itoa(i))
		rabbitmqTwo.PublishTopic("hello huxiaobai two" + strconv.Itoa(i))
		time.Sleep(1 * time.Second / 10)
		fmt.Println(i)
	}
}
