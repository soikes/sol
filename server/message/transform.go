package message

import (
	"encoding/json"

	"soikke.li/sol/primitives"
)

type Transform struct {
	ID  string          `json:"id"`
	Pos primitives.Vec3 `json:"pos"`
	Rot primitives.Vec3 `json:"rot"`
}

func (t *Transform) Marshal() ([]byte, error) {
	e := Msg{
		Type: MsgTransform,
		Data: t,
	}
	return json.Marshal(e)
}
