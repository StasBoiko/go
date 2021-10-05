package main

import (
	"encoding/json"
	"encoding/xml"
)

type Validater interface {
	Validate(string) error
}

type Json struct{}
type Xml struct{}

func NewValidator(flag string) Validater {
	switch flag {
	case "json":
		return &Json{}
	case "xml":
		return &Xml{}
	}
	return nil
}

func (j *Json) Validate(text string) error {
	var js jsonData
	return json.Unmarshal([]byte(text), &js)
}

func (j *Xml) Validate(text string) error {
	var x interface{}
	return xml.Unmarshal([]byte(text), &x)
}
