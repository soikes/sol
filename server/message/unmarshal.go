package message

import (
	"encoding/json"
	"fmt"
)

func Unmarshal(b []byte) (Input, error) { //TODO how can we return any kind of message type
	var d json.RawMessage
	e := Envelope{
		Data: &d,
	}
	if err := json.Unmarshal(b, &e); err != nil {
		return Input{}, err
	}
	switch e.Type {
	case TypeInput:
		var i Input
		if err := json.Unmarshal(d, &i); err != nil {
			return Input{}, err
		}
		return i, nil
	}
	return Input{}, fmt.Errorf(`invalid message type: %d`, e.Type)
}