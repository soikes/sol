package message

type MsgType int

const (
	MsgInput MsgType = iota
	MsgTransform
	MsgSpawn
	MsgRegister
)

type SpawnType int

const (
	SpawnPlayer SpawnType = iota
)
