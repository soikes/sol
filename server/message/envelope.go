//package Message describes input and output message formats that can be used to communicate between sol clients and a sol server.
package message

const (
	TypeInput = iota
)

type Envelope struct {
	Type int	`json:"type"`
	Data interface{} `json:"data"`
}

