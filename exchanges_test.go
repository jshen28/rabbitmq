package rabbitmq

import (
	"fmt"
	"log"
)

func ExampleClient_Exchanges() {
	client, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}

	exchanges, err := client.Exchanges().Do()
	if err != nil {
		log.Fatal(err)
	}

	if len(exchanges) == 0 {
		fmt.Printf("No exchanges found")
		return
	}

	fmt.Printf("Success")

	// Output:
	// Success
}
