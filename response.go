package main

import (
	"encoding/json"
	"fmt"
)

type Return struct {
	Version           string            `json:"version"`
	SessionAttributes SessionAttributes `json:"SessionAttributes"`
	Response          Response          `json:"response"`
	ShouldEndSession  bool              `json:"shouldEndSession"`
}

func NewReturn(s SessionAttributes, r Response, e bool) Return {
	return Return{
		Version:           "1.0",
		SessionAttributes: s,
		Response:          r,
		ShouldEndSession:  e,
	}
}

func PrintReturn(r Return) error {
	j, err := json.Marshal(r)
	fmt.Print(string(j))
	return err
}

type SessionAttributes map[string]interface{}

type Response struct {
	OutputSpeech PlainText `json:"outputSpeech"`
	Card         Card      `json:"card"`
	Reprompt     PlainText `json:"reprompt"`
}

func NewResponse(o string, t string, c string, r string) Response {
	return Response{
		NewPlainText(o),
		NewSimpleCard(t, c),
		NewPlainText(r),
	}
}

type PlainText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewPlainText(t string) PlainText {
	return PlainText{
		Type: "PlainText",
		Text: t,
	}
}

type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewSimpleCard(t string, c string) Card {
	return Card{
		Type:    "Simple",
		Title:   t,
		Content: c,
	}
}
