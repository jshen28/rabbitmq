package rabbitmq

import (
	"fmt"
	"log"
)

func ExampleClient_Queues() {
	client, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}

	queues, err := client.Queues().Do()
	if err != nil {
		log.Fatal(err)
	}

	if len(queues) == 0 {
		fmt.Printf("No queues found")
		return
	}

	fmt.Printf("Success")

	// Output:
	// Success
}
