package posbus

import (
	"encoding/binary"
	"math"

	"github.com/momentum-xyz/controller/pkg/cmath"
	"github.com/momentum-xyz/posbus-protocol/utils"

	"github.com/google/uuid"
)

const TriggerTransitionalBridgingEffectsOnPositionElementSize = MsgUUIDTypeSize + 6*MsgOnePosSize + MsgTypeSize

type TriggerTransitionalBridgingEffectsOnPosition struct {
	*Message
}

func NewTriggerTransitionalBridgingEffectsOnPositionMsg(numEffects int) *TriggerTransitionalBridgingEffectsOnPosition {
	obj := NewMessage(
		MsgTypeTriggerTransitionalBridgingEffectsOnPosition,
		MsgArrTypeSize+numEffects*TriggerTransitionalBridgingEffectsOnPositionElementSize,
	)
	binary.LittleEndian.PutUint32(obj.Msg(), uint32(numEffects))
	return &TriggerTransitionalBridgingEffectsOnPosition{
		Message: obj,
	}
}

func (m *TriggerTransitionalBridgingEffectsOnPosition) SetEffect(
	i int, emitter uuid.UUID, from, to cmath.Vec3, effect uint32,
) {
	start := MsgArrTypeSize + i*TriggerTransitionalBridgingEffectsOnPositionElementSize
	copy(m.Msg()[start:], utils.BinID(emitter))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize:], math.Float32bits(from.X))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize+MsgOnePosSize:], math.Float32bits(from.Y))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize+2*MsgOnePosSize:], math.Float32bits(from.Z))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize+3*MsgOnePosSize:], math.Float32bits(to.X))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize+4*MsgOnePosSize:], math.Float32bits(to.Y))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize+5*MsgOnePosSize:], math.Float32bits(to.Z))
	binary.LittleEndian.PutUint32(m.Msg()[start+MsgUUIDTypeSize+6*MsgOnePosSize:], effect)
}
