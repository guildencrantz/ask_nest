package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/otherinbox/gobuild"
)

type Session struct {
	New bool `json:"new"`
}

type Intent struct {
	Name  string                            `json:"name"`
	Slots map[string]map[string]interface{} `json:"slots"`
}

type Request struct {
	Type   string `json:"type"`
	Id     string `json:"requestId"`
	Intent Intent `json:"intent"`
}

type Event struct {
	Session Session `json:"session"`
	Request Request `json:"request"`
}

func Version() string {
	return gobuild.Version
}

func main() {
	input := os.Args[1]
	fmt.Println(input)

	var event Event
	json.Unmarshal([]byte(input), &event)
	fmt.Println("%q", event)
	fmt.Println(event.Request.Intent.Slots["Color"]["value"])
}
