package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/otherinbox/gobuild"
)

const (
	LaunchRequest = "LaunchRequest"
	IntentRequest = "IntentRequest"

	HelpIntent        = "HelpIntent"
	StatusIntent      = "StatusIntent"
	SetTempIntent     = "SetTempIntent"
	SetPresenceIntent = "SetPresenceIntent"
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

	var event Event
	json.Unmarshal([]byte(input), &event)
	fmt.Fprintln(os.Stderr, string(input))

	var r Return
	switch event.Request.Type {
	case LaunchRequest:
		r = Help()
	case IntentRequest:
		switch event.Request.Intent.Name {
		case HelpIntent:
			r = Help()
		case StatusIntent:
			r = Status()
		case SetTempIntent:
			r = SetTemp(event.Request.Intent)
		case SetPresenceIntent:
			r = SetPresence(event.Request.Intent)
		default:
			// TODO: Special case help.
			r = Unknown()
		}
	}

	if err := PrintReturn(r); err != nil {
		os.Exit(1)
	}
}

func Unknown() Return {
	return NewReturn(
		nil,
		NewResponse(
			"I don't understand your request",
			"I don't understand.",
			"I should try to help you",
			"",
		),
		true,
	)
}

func Help() Return {
	return NewReturn(
		nil,
		NewResponse("Help", "Help", "", ""),
		true,
	)
}

func Status() Return {
	return NewReturn(
		nil,
		NewResponse("Status", "Status", "", ""),
		true,
	)
}

func SetTemp(i Intent) Return {
	return NewReturn(
		nil,
		NewResponse("Set Temp", "Set Temp", "", ""),
		true,
	)
}

func SetPresence(i Intent) Return {
	return NewReturn(
		nil,
		NewResponse("Set Presence", "Set Presence", "", ""),
		true,
	)
}
