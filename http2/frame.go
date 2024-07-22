package http2

type Frame struct {
	length    uint32 // u24: the first 8 bits of the frame MUST be disregarded
	ttype     uint8
	flags     uint8
	stream_id uint32 // u31: TODO: servers MUST disregard the most significant bit. clients MUST set to 0
	payload   []byte
}

type FrameType int

const (
	Data FrameType = iota
	Headers
	Priority
	RstStream
	Settings
	PushPromise
	Ping
	Goaway
	WindowUpdate
	Continuation
)

const frameName map[FrameType]string = map[FrameType]string{
	Data:         "DATA",
	Headers:      "HEADERS",
	Priority:     "PRIORITY",
	RstStream:    "RST_STREAM",
	Settings:     "SETTINGS",
	PushPromise:  "PUSH_PROMISE",
	Ping:         "PING",
	Goaway:       "GOAWAY",
	WindowUpdate: "WINDOW_UPDATE",
	Continuation: "CONTINUATION",
}

type SettingsFlag int

const (
	SettingsHeaderTableSize      SettingsFlag = iota + 1
	SettingsEnablePush
	SettingsMaxConcurrentStreams
	SettingsInitialWindowSize
	SettingsMaxFrameSize
	SettingsMaxHeaderListSize
)

const settingsFlagName map[SettingsFlag]string = map[SettingsFlag]string{
	SettingsHeaderTableSize:      "SETTINGS_HEADER_TABLE_SIZE",
	SettingsEnablePush:           "SETTINGS_ENABLE_PUSH",
	SettingsMaxConcurrentStreams: "SETTINGS_MAX_CONCURRENT_STREAMS",
	SettingsInitialWindowSize:    "SETTINGS_INITIAL_WINDOW_SIZE",
	SettingsMaxFrameSize:         "SETTINGS_MAX_FRAME_SIZE",
	SettingsMaxHeaderListSize:    "SETTINGS_MAX_HEADER_LIST_SIZE",
}

type FrameErrorCode int32

// TODO: when receiving an unknown error code, treat them as InternalError
const (
	NoError FrameErrorCode = iota
	ProtocolError
	InternalError
	FlowControlError
	SettingsTimeout
	StreamClosed
	FrameSizeError
	RefusedStream
	Cancel
	CompressionError
	ConnectError
	EnhanceYourCalm
	InadequateSecurity
	Http11Required
)

func (t FrameType) String() string {
	name, ok := frameName[t]
	if !ok {
		return "UNKNOWN_FRAME_TYPE"
	}
	return name
}

func (f SettingsFlag) String() string {
	name, ok := settingsFlagName[f]
	if !ok {
		return "UNKNOWN_SETTINGS_FLAG"
	}
	return name
}

func (f *Frame) Serialize() []byte {
	return []byte{
		byte(f.length >> 16),
		byte(f.length >> 8),
		byte(f.length),
		byte(f.ttype),
		byte(f.flags),
		byte(f.stream_id >> 24),
		byte(f.stream_id >> 16),
		byte(f.stream_id >> 8),
		byte(f.stream_id),
	}
}

func ParseFrame(data []byte) Frame {
	return Frame{
		length:    uint32(data[0])<<16 | uint32(data[1])<<8 | uint32(data[2]),
		ttype:     data[3],
		flags:     data[4],
		stream_id: uint32(data[5])<<24 | uint32(data[6])<<16 | uint32(data[7])<<8 | uint32(data[8]),
		payload:   data[9:],
}
