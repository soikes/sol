package message

type MsgType int

const (
	MsgInput MsgType = iota
	MsgTransform
	MsgSpawn
	MsgRegister
	MsgSync
)

type SpawnType int

const (
	SpawnPlayer SpawnType = iota
	SpawnOtherPlayer
)
