package message

import "encoding/json"

type Sync struct {
}

func (s *Sync) Marshal() ([]byte, error) {
	e := Msg{
		Type: MsgSync,
		Data: s,
	}
	return json.Marshal(e)
}
