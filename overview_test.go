package rabbitmq

import (
	"fmt"
	"log"
)

func ExampleClient_Overview() {
	client, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}

	overview, err := client.Overview().Do()
	if err != nil {
		log.Fatal(err)
	}

	if overview.RabbitMQVersion == "" {
		fmt.Printf("RabbitMQVersion is not set")
		return
	}

	fmt.Printf("Success")

	// Output:
	// Success
}
