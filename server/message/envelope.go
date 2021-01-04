//package message describes input and output message formats that can be used to communicate between sol clients and a sol server.
package message

type Envelope struct {
	Type MsgType     `json:"type"`
	Data interface{} `json:"data"`
}
