package session

import (
	"errors"
	simplefixgo "github.com/b2broker/simplefix-go"
)

var (
	ErrNotEnoughMessages = errors.New("Not enough messages in the storage")
	ErrInvalidBoundaries = errors.New("Invalid boundaries")
	ErrInvalidSequence   = errors.New("Unexpected sequence index")
)

// MessageStorage is an interface providing a basic method for storing messages awaiting to be sent.
type MessageStorage interface {
	Save(msg simplefixgo.SendingMessage, msgSeqNum int) error
	Messages(msgSeqNumFrom, msgSeqNumTo int) ([]simplefixgo.SendingMessage, error)
}
