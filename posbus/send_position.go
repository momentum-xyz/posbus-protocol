package posbus

import (
	"encoding/binary"
	"math"

	"github.com/momentum-xyz/controller/pkg/cmath"
)

type SendPosition struct {
	*Message
}

func NewSendPositionMsg(pos cmath.Vec3) *SendPosition {
	obj := NewMessage(MsgTypeSendPosition, 3*MsgOnePosSize)
	nobj := SendPosition{
		Message: obj,
	}
	nobj.SetPosition(pos)
	return &nobj
}

func (m *SendPosition) SetPosition(pos cmath.Vec3) {
	binary.LittleEndian.PutUint32(m.Msg(), math.Float32bits(pos.X))
	binary.LittleEndian.PutUint32(m.Msg()[MsgOnePosSize:], math.Float32bits(pos.Y))
	binary.LittleEndian.PutUint32(m.Msg()[2*MsgOnePosSize:], math.Float32bits(pos.Z))
}
