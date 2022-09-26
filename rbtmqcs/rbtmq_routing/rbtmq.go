package rbtmq_routing

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQURl = "amqp://huj:123456@127.0.0.1:5672/rbtmq"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

func NewRabbitMQ(queuename string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queuename,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURl,
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误！")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败！")
	return rabbitmq
}
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, routingKey)
}

func (r *RabbitMQ) PublishRouting(message string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil)
	r.failOnErr(err, "Failed to declare an exchange")

	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
func (r *RabbitMQ) ConsumeRouting() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil)
	r.failOnErr(err, "Failed to declare an exchange")

	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil)
	r.failOnErr(err, "Failed to declare a queue")

	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil)

	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Recieved a message : %s", d.Body)
		}
		log.Printf("[*] Waiting for message, To exit pree CTRL+C")
	}()
	<-forever
}
