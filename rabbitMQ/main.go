package main

import (
	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672")

	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("first_queue", true, false, false, false, nil)
	//ch.QueueBind(q.Name, "","amq.fanout",false,nil)

	msg := amqp.Publishing{
		Body: []byte("first message"),
	}

	ch.Publish("", q.Name, false, false, msg)
	// exchange can be amq.fanout

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false,nil)

	for m := range msgs {
		println(string(m.Body))
	}

}
