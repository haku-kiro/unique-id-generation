package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func getId() (int, error) {
	resp, err := http.Get("http://localhost:8090/serve-id")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// The id server should return just a single number, no other content.
	scanner := bufio.NewScanner(resp.Body)
	scanner.Scan()
	data := scanner.Text()
	return strconv.Atoi(data)
}

func main() {
	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		id, err := getId()
		if err != nil {
			return err
		}

		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "generate_ok"
		body["id"] = id

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
