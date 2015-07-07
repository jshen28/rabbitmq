package rabbitmq

import (
	"fmt"
	"log"
)

func ExampleClient_QueueByName() {
	client, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.QueueByName("/", "no.such.queue.name").Do()
	if err == nil {
		log.Println("expected no such queue")
	}

	got, want := err.Error(), `rabbitmq: Error 404 (Not Found): Object Not Found ("Not Found")`
	if want != got {
		fmt.Printf("expected error %q; got: %q", want, got)
		return
	}

	fmt.Printf("Success")

	// Output:
	// Success
}
