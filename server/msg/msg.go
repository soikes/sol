package msg

import "encoding/json"

type Msg struct {
	Type MsgType         `json:"type"`
	Data json.RawMessage `json:"data"`
}

// type UserInputMsg struct {
// 	Mt   MsgType   `json:"mt"`
// 	Data UserInput `json:"data"`
// }

type RegisterMsg struct {
	ClientID string `json:"client_id"`
	Error    string `json:"error"`
}
