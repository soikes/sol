package message

type Input struct {
	ID string `json:"id"`

	ForwardPress bool `json:"forward"`
	BackwardPress bool `json:"backward"`
	LeftPress bool `json:"left"`
	RightPress bool `json:"right"`
	FirePress bool `json:"fire"`
}