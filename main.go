package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var (
	counter uint64
	mu      sync.Mutex
)

func generateUniqueID() string {
	mu.Lock()
	defer mu.Unlock()
	counter++
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), counter)
}

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		// Generate unique ID
		uniqueID := generateUniqueID()

		// Create the response
		response := map[string]interface{}{
			"type": "generate_ok",
			"id":   uniqueID,
		}

		// Send the response
		return n.Reply(msg, response)

	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
