package messages

type ResendRequestBuilder interface {
	// Flow Message
	Parse(data []byte) (ResendRequestBuilder, error)
	New() ResendRequestBuilder

	// Generated ResendRequest Message
	BeginSeqNo() int
	SetFieldBeginSeqNo(int) ResendRequestBuilder
	EndSeqNo() int
	SetFieldEndSeqNo(int) ResendRequestBuilder

	// HeaderBuilder
	HeaderBuilder() HeaderBuilder

	// SendingMessage
	MsgType() string
	ToBytes() ([]byte, error)
}