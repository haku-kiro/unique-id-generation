package main

import (
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
	defer client.Close()

	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		id, err := idCall(client)
		if err != nil {
			return err
		}

		response := map[string]any{
			"type": "generate_ok",
			"id":   id,
		}

		return n.Reply(msg, response)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
