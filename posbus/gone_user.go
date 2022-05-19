package posbus

import (
	"encoding/binary"

	// Momentum
	"github.com/momentum-xyz/posbus-protocol/utils"

	// Third-Party
	"github.com/google/uuid"
)

const GoneUsersElementSize = MsgUUIDTypeSize

type GoneUsersMessage struct {
	*Message
}

func NewGoneUsersMsg(numUsers int) *GoneUsersMessage {
	obj := NewMessage(MsgTypeGoneUsers, MsgArrTypeSize+numUsers*GoneUsersElementSize)
	binary.LittleEndian.PutUint32(obj.Msg(), uint32(numUsers))
	return &GoneUsersMessage{
		Message: obj,
	}
}

func (m *GoneUsersMessage) SetUser(i int, id uuid.UUID) {
	start := MsgArrTypeSize + i*GoneUsersElementSize
	copy(m.Msg()[start:], utils.BinID(id))
}
