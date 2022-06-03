package posbus

import (
	"encoding/binary"
	"math"

	"github.com/momentum-xyz/controller/pkg/cmath"
	"github.com/momentum-xyz/posbus-protocol/utils"

	"github.com/google/uuid"
)

const TriggerTransitionalEffectsOnPositionElementSize = MsgUUIDTypeSize + 3*MsgOnePosSize + MsgTypeSize

type TriggerTransitionalEffectsOnPosition struct {
	*Message
}

func NewTriggerTransitionalEffectsOnPositionMsg(numEffects int) *TriggerTransitionalEffectsOnPosition {
	obj := NewMessage(
		MsgTypeTriggerTransitionalEffectsOnPosition,
		MsgArrTypeSize+numEffects*TriggerTransitionalEffectsOnPositionElementSize,
	)
	binary.LittleEndian.PutUint32(obj.Msg(), uint32(numEffects))
	return &TriggerTransitionalEffectsOnPosition{
		Message: obj,
	}
}

func (m *TriggerTransitionalEffectsOnPosition) SetEffect(i int, emitter uuid.UUID, pos cmath.Vec3, effect uint32) {
	start := MsgArrTypeSize + i*TriggerTransitionalEffectsOnPositionElementSize
	copy(m.Msg()[start:], utils.BinID(emitter))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize:], math.Float32bits(pos.X))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize+MsgOnePosSize:], math.Float32bits(pos.Y))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize+2*MsgOnePosSize:], math.Float32bits(pos.Z))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize+3*MsgOnePosSize:], effect)
}
