package message

import (
	"encoding/json"
	"fmt"
)

func Unmarshal(b []byte, in interface{}) error {
	var d json.RawMessage
	e := Envelope{
		Data: &d,
	}
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}
	switch e.Type {
	case MsgInput:
		if err := json.Unmarshal(d, in); err != nil {
			return err
		}
	}
	return fmt.Errorf(`invalid message type: %d`, e.Type)
}

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/rs/zerolog/log"
// )

// type Marshaler struct {
// 	incoming chan []byte
// 	outgoing chan []byte

// 	inputsIn chan Input

// 	transformsOut chan Transform
// }

// func (m *Marshaler) inPump() {

// 	var d json.RawMessage
// 	e := Envelope{
// 		Data: &d,
// 	}
// 	if err := json.Unmarshal(b, &e); err != nil {
// 		return err
// 	}
// 	switch e.Type {
// 	case InputMsg:
// 		var i Input
// 		if err := json.Unmarshal(d, &i); err != nil {
// 			return err
// 		}
// 		m.inputsIn <- i
// 		return nil
// 	}
// 	return fmt.Errorf(`invalid message type: %d`, e.Type)
// }

// func (m *Marshaler) outPump() {
// 	var e Envelope
// 	for {
// 		select {
// 		case t := <-m.transformsOut:
// 			e.Type = TransformMsg
// 			e.Data = t
// 		}
// 		out, err := json.Marshal(e)
// 		if err != nil {
// 			log.Warn().Err(err).Msg(`failed to marshal outbound message`)
// 		}
// 		m.outgoing <- out
// 	}
// }

// func (m *Marshaler) Start() {
// 	go m.inPump()
// 	go m.outPump()
// }
