package posbus

import (
	"encoding/binary"
	"math"

	"github.com/momentum-xyz/posbus-protocol/utils"

	"github.com/google/uuid"
)

type SetStaticObjectPosition struct {
	*Message
}

func NewSetStaticObjectPositionMsg() *SetStaticObjectPosition {
	obj := NewMessage(MsgTypeSetStaticObjectPosition, UserPositionsMessageSize)
	return &SetStaticObjectPosition{
		Message: obj,
	}
}

func (m *SetStaticObjectPosition) SetPosition(id uuid.UUID, pos utils.Vec3) {
	copy(m.Msg()[:MsgUUIDTypeSize], utils.BinID(id))
	binary.LittleEndian.PutUint32(m.Msg()[MsgUUIDTypeSize:], math.Float32bits(pos.X))
	binary.LittleEndian.PutUint32(m.Msg()[MsgUUIDTypeSize+MsgOnePosSize:], math.Float32bits(pos.Y))
	binary.LittleEndian.PutUint32(m.Msg()[MsgUUIDTypeSize+2*MsgOnePosSize:], math.Float32bits(pos.Z))
}
