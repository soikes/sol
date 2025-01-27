package message

import "encoding/json"

type Register struct {
	ID string `json:"id"`
}

func (r *Register) Marshal() ([]byte, error) {
	e := Msg{
		Type: MsgRegister,
		Data: r,
	}
	return json.Marshal(e)
}
