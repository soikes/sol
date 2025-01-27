package msg

type UserInput int

const (
	UserInputForwardStart UserInput = iota
	UserInputForwardEnd
	UserInputBackwardStart
	UserInputBackwardEnd
	UserInputTurnLeftStart
	UserInputTurnLeftEnd
	UserInputTurnRightStart
	UserInputTurnRightEnd
)
