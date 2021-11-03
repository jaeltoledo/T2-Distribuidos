package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("consumer app")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("Received: %s\n", d.Body)

		}
	}()

	fmt.Println(("successfull"))
	fmt.Println(" [*] waiting")
	<-forever

}