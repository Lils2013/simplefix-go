package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
	"github.com/b2broker/simplefix-go/session/messages"
)

const MsgTypeMarketDataIncrementalRefresh = "X"

type MarketDataIncrementalRefresh struct {
	*fix.Message
}

func makeMarketDataIncrementalRefresh() *MarketDataIncrementalRefresh {
	msg := &MarketDataIncrementalRefresh{
		Message: fix.NewMessage(FieldBeginString, FieldBodyLength, FieldCheckSum, FieldMsgType, beginString, MsgTypeMarketDataIncrementalRefresh).
			SetBody(
				fix.NewKeyValue(FieldMDReqID, &fix.String{}),
				NewMDEntriesGrp().Group,
				fix.NewKeyValue(FieldApplQueueDepth, &fix.Int{}),
				fix.NewKeyValue(FieldApplQueueResolution, &fix.String{}),
			),
	}

	msg.SetHeader(makeHeader().AsComponent())
	msg.SetTrailer(makeTrailer().AsComponent())

	return msg
}

func NewMarketDataIncrementalRefresh(header *Header, trailer *Trailer, noMDEntries *MDEntriesGrp) *MarketDataIncrementalRefresh {
	msg := makeMarketDataIncrementalRefresh().
		SetMDEntriesGrp(noMDEntries)
	msg.SetHeader(header.AsComponent())
	msg.SetTrailer(trailer.AsComponent())
	return msg
}

func ParseMarketDataIncrementalRefresh(data []byte) (*MarketDataIncrementalRefresh, error) {
	msg := fix.NewMessage(FieldBeginString, FieldBodyLength, FieldCheckSum, FieldMsgType, FieldBeginString, beginString).
		SetBody(makeMarketDataIncrementalRefresh().Body()...).
		SetHeader(makeHeader().AsComponent()).
		SetTrailer(makeTrailer().AsComponent())

	if err := msg.Unmarshal(data); err != nil {
		return nil, err
	}

	return &MarketDataIncrementalRefresh{
		Message: msg,
	}, nil
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) Header() *Header {
	header := marketDataIncrementalRefresh.Message.Header()

	return &Header{header}
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) HeaderBuilder() messages.HeaderBuilder {
	return marketDataIncrementalRefresh.Header()
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) Trailer() *Trailer {
	trailer := marketDataIncrementalRefresh.Message.Trailer()

	return &Trailer{trailer}
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) MDReqID() string {
	kv := marketDataIncrementalRefresh.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) SetMDReqID(mDReqID string) *MarketDataIncrementalRefresh {
	kv := marketDataIncrementalRefresh.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(mDReqID)
	return marketDataIncrementalRefresh
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) MDEntriesGrp() *MDEntriesGrp {
	group := marketDataIncrementalRefresh.Get(1).(*fix.Group)

	return &MDEntriesGrp{group}
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) SetMDEntriesGrp(noMDEntries *MDEntriesGrp) *MarketDataIncrementalRefresh {
	marketDataIncrementalRefresh.Set(1, noMDEntries.Group)

	return marketDataIncrementalRefresh
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) ApplQueueDepth() int {
	kv := marketDataIncrementalRefresh.Get(2)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) SetApplQueueDepth(applQueueDepth int) *MarketDataIncrementalRefresh {
	kv := marketDataIncrementalRefresh.Get(2).(*fix.KeyValue)
	_ = kv.Load().Set(applQueueDepth)
	return marketDataIncrementalRefresh
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) ApplQueueResolution() string {
	kv := marketDataIncrementalRefresh.Get(3)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (marketDataIncrementalRefresh *MarketDataIncrementalRefresh) SetApplQueueResolution(applQueueResolution string) *MarketDataIncrementalRefresh {
	kv := marketDataIncrementalRefresh.Get(3).(*fix.KeyValue)
	_ = kv.Load().Set(applQueueResolution)
	return marketDataIncrementalRefresh
}