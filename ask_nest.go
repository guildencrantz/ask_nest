package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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
	Name  string                       `json:"name"`
	Slots map[string]map[string]string `json:"slots"`
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
	r := fmt.Sprintf("The current temperature is %d", thermostat.AmbientTemperatureF)
	return NewReturn(
		nil,
		NewResponse(
			r,
			"Status",
			r,
			"",
		),
		true,
	)
}

func SetTemp(i Intent) Return {
	// This wouldn't work if I used nest.HeatCool
	ct := thermostat.TargetTemperatureF
	// TODO: Create an error response condition.
	tt, _ := strconv.Atoi(i.Slots["Temperature"]["value"])

	if tt == ct {
		r := fmt.Sprintf("You requested the target temperature be set to %d, but that's already the target temperature.", tt)
		return NewReturn(nil, NewResponse(r, "Target Temperature not changed.", r, ""), true)
	}

	// Need that error response stuff here too.
	thermostat.SetTargetTempF(tt)

	a := thermostat.AmbientTemperatureF
	h := thermostat.HvacMode

	fmt.Fprintf(os.Stderr, "h: %#v", h)

	var r string
	if a == tt {
		r = fmt.Sprintf("Holding temperature at %d", tt)
	} else if a > tt {
		switch h {
		case "heat":
			r = fmt.Sprintf("Will allow temperature to drop from %d to %d", a, tt)
		default:
			r = fmt.Sprintf("Cooling from %d to %d", a, tt)
		}
	} else {
		switch h {
		case "heat":
			r = fmt.Sprintf("Heating from %d to %d", a, tt)
		default:
			r = fmt.Sprintf("Will allow temperature to rise from %d to %d", a, tt)
		}
	}

	return NewReturn(
		nil,
		NewResponse(r, "Set Temperature", r, ""),
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
