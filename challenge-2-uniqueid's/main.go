package main

import (
	"log"

	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any = make(map[string]any)

		id := uuid.New().String()
		body["type"] = "generate_ok"
		body["id"] = id
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
