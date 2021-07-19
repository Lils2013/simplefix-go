package messages

type SequenceResetBuilder interface {
	// Flow Message
	Parse(data []byte) (SequenceResetBuilder, error)
	New() SequenceResetBuilder

	// Generated ResendRequest Message
	NewSeqNo() int
	SetFieldNewSeqNo(newSeqNo int) SequenceResetBuilder

	// HeaderBuilder
	HeaderBuilder() HeaderBuilder

	// SendingMessage
	MsgType() string
	ToBytes() ([]byte, error)
}