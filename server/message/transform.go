package message

import "encoding/json"

type Transform struct {
	ID string
	X  float64
	Y  float64
	Z  float64
}

func (t *Transform) Marshal() ([]byte, error) {
	e := Envelope{
		Type: MsgTransform,
		Data: t,
	}
	return json.Marshal(e)
}
