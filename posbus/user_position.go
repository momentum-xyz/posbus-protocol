package posbus

import (
	"encoding/binary"
)

const UserPositionsMessageSize = MsgUUIDTypeSize + MsgOnePosSize*3

type UserPositionsMessage struct {
	*Message
}

func NewUserPositionsMsg(numUsers int) *UserPositionsMessage {
	obj := NewMessage(MsgTypeUsersPositions, MsgArrTypeSize+numUsers*UserPositionsMessageSize)
	binary.LittleEndian.PutUint32(obj.Msg(), uint32(numUsers))
	return &UserPositionsMessage{
		Message: obj,
	}
}

func (m *UserPositionsMessage) SetPosition(i int, data []byte) {
	start := MsgArrTypeSize + i*UserPositionsMessageSize
	copy(m.Msg()[start:], data)
}
