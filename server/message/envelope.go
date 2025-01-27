// package message describes input and output message formats that can be used to communicate between sol clients and a sol server.
package message

import "time"

type Msg struct {
	Type      MsgType   `json:"type"`
	Timestamp time.Time `json:"ts"`
	Data      any       `json:"data"`
}
