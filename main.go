package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var (
	filePath = "./id"
)

func storeId(id int) error {
	stringId := fmt.Sprint(id)
	// Assuming we're overwriting the file each time?
	return os.WriteFile(filePath, []byte(stringId), 0644)
}

func readId() (int, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(data))
}

func main() {
	// Seeding our id
	err := storeId(0)
	if err != nil {
		log.Fatal(err)
	}

	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		id, err := readId()
		if err != nil {
			return err
		}
		id += 1

		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "generate_ok"
		body["id"] = id

		err = storeId(id)
		if err != nil {
			return err
		}

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
