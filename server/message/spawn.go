package message

import "encoding/json"

type Spawn struct {
	ID string `json:"id"`

	Type SpawnType `json:"type"`
}

func (s *Spawn) Marshal() ([]byte, error) {
	e := Envelope{
		Type: MsgSpawn,
		Data: s,
	}
	return json.Marshal(e)
}
