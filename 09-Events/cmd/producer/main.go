package main

import "github.com/AndreCDiniz/PosGoExpert/fcutils/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, World from Producer!", "amq.direct")
}
