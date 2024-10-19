package main

import (
	"encoding/json"
	"log"
	"net/rpc"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	gen "unique.ids/internal/id-gen"
)

func idCall(c *rpc.Client) (uint64, error) {
	req := gen.Request{}
	resp := new(gen.Response)

	incrementCall := c.Go("Inc.Next", req, resp, nil)
	reply := <-incrementCall.Done

	if reply.Error != nil {
		log.Fatalf("reply: %d, reply error: %s", resp.Id, reply.Error.Error())
	}

	return resp.Id, nil
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		id, err := idCall(client)
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
