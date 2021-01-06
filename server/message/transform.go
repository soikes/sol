package message

import "encoding/json"

type Transform struct {
	ID string  `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	Z  float64 `json:"z"`
}

func (t *Transform) Marshal() ([]byte, error) {
	e := Envelope{
		Type: MsgTransform,
		Data: t,
	}
	return json.Marshal(e)
}
