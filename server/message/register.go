package message

import "encoding/json"

type Register struct {
	ID string
}

func (r *Register) Marshal() ([]byte, error) {
	e := Envelope{
		Type: MsgRegister,
		Data: r,
	}
	return json.Marshal(e)
}
