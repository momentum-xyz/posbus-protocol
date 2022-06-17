package posbus

import (
	"encoding/binary"

	"github.com/gorilla/websocket"

	"github.com/momentum-xyz/posbus-protocol/flatbuff/go/api"
)

const (
	MsgTypeSize     = 4
	MsgArrTypeSize  = 4
	MsgUUIDTypeSize = 16
	MsgOnePosSize   = 4
)

type MsgType uint32

/* can use fmt.Sprintf("%x", int) to display hex */
const (
	MsgTypeNONE                                         MsgType = 0x00000000
	MsgTypeFlatBufferMessage                            MsgType = 0x49a0b0b9
	MsgTypeUsersPositions                               MsgType = 0x1FE5B46F // 535147631 (uint32)
	MsgTypeSendPosition                                 MsgType = 0xfbf6b89f
	MsgTypeRemoveStaticObjects                          MsgType = 0x06383502
	MsgTypeSetStaticObjectPosition                      MsgType = 0x300A3883
	MsgTypeTriggerTransitionalEffectsOnObject           MsgType = 0xE0A9E0A7
	MsgTypeTriggerTransitionalEffectsOnPosition         MsgType = 0x3597729E
	MsgTypeTriggerTransitionalBridgingEffectsOnObject   MsgType = 0xE45A7B03
	MsgTypeTriggerTransitionalBridgingEffectsOnPosition MsgType = 0xF6AB754D
	MsgTypeSignal                                       MsgType = 0x6A8634A3
	MsgTypeSwitchWorld                                  MsgType = 0x7D40FD67
	MsgTriggerInteraction                               MsgType = 0x2C0A16A0
	MsgTypeGoneUsers                                    MsgType = 0x3327c20c
	MsgTypeSimpleNotification                           MsgType = 0x3CADFD52
	MsgTypeNotificationSingleID                         MsgType = 0x91D2F62B
	MsgTypeNotificationDualID                           MsgType = 0xC66E8C60
	MsgTypeRelayToReact                                 MsgType = 0xB5BBCFA2
)

type Signal uint32

const (
	SignalNone Signal = iota
	SignalDualConnection
	SignalReady
	SignalInvalidToken
)

type Trigger uint32

const (
	TriggerNone = iota
	TriggerWow
	TriggerHighFive
	TriggerEnteredSpace
	TriggerLeftSpace
	TriggerStake
)

type Notification uint32

const (
	NotificationNone     Notification = 0
	NotificationWow      Notification = 1
	NotificationHighFive Notification = 2

	NotificationStageModeAccept        Notification = 10
	NotificationStageModeInvitation    Notification = 11
	NotificationStageModeSet           Notification = 12
	NotificationStageModeStageJoin     Notification = 13
	NotificationStageModeStageRequest  Notification = 14
	NotificationStageModeStageDeclined Notification = 15

	NotificationTextMessage Notification = 500
	NotificationRelay       Notification = 501

	NotificationGeneric Notification = 999
	NotificationLegacy  Notification = 1000
)

type Destination byte

const (
	DestinationUnity Destination = 0b01
	DestinationReact Destination = 0b10
	DestinationBoth  Destination = 0b11
)

type Message struct {
	buf []byte
}

func NewMessage(msgid MsgType, len int) *Message {
	obj := &Message{}
	obj.Allocate(len)
	binary.LittleEndian.PutUint32(obj.buf, uint32(msgid))
	binary.LittleEndian.PutUint32(obj.buf[MsgTypeSize+len:], uint32(^msgid))
	return obj
}

func (m *Message) Buf() []byte {
	return m.buf
}

func (m *Message) Msg() []byte {
	return m.buf[MsgTypeSize : len(m.buf)-MsgTypeSize]
}

func (m *Message) Type() MsgType {
	/*if len(m.buf) < 4 {
		// TODO: Handle
	}*/
	header := binary.LittleEndian.Uint32(m.buf[:MsgTypeSize])
	footer := binary.LittleEndian.Uint32(m.buf[len(m.buf)-MsgTypeSize:])
	if header == ^footer {
		return MsgType(header)
	}
	return MsgTypeNONE
}

func (m *Message) AsFlatBufferMessage() *api.FlatBuffMsg {
	if m.Type() != MsgTypeFlatBufferMessage {
		return nil
	}

	return api.GetRootAsFlatBuffMsg(m.buf, MsgTypeSize)
}

func (m *Message) AsSendPos() []byte {
	// how 16 have been calculated?
	return m.buf[MsgTypeSize:16]
}

func (m *Message) Allocate(len int) {
	m.buf = make([]byte, MsgTypeSize*2+len)
}

func (m *Message) WebsocketMessage() *websocket.PreparedMessage {
	omsg, _ := websocket.NewPreparedMessage(websocket.BinaryMessage, m.Buf())
	return omsg
}

func MsgFromBytes(b []byte) *Message {
	return &Message{
		buf: b,
	}
}
