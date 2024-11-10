package models

import "encoding/json"

type Ops struct {
	Attributes struct {
		Header     int    `json:"header,omitempty"`
		Bold       bool   `json:"bold,omitempty"`
		Italic     bool   `json:"italic,omitempty"`
		Underline  bool   `json:"underline,omitempty"`
		Strike     bool   `json:"strike,omitempty"`
		Blockquote bool   `json:"blockquote,omitempty"`
		CodeBlock  string `json:"code-block,omitempty"`
		Link       string `json:"link,omitempty"`
		Image      string `json:"image,omitempty"`
	} `json:"attributes,omitempty"`
	Insert json.RawMessage `json:"insert,omitempty"`
}

func NewOps() *Ops {
	return &Ops{}
}
