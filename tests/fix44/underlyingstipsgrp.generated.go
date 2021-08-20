package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
)

type UnderlyingStipsGrp struct {
	*fix.Group
}

func NewUnderlyingStipsGrp() *UnderlyingStipsGrp {
	return &UnderlyingStipsGrp{
		fix.NewGroup(FieldNoUnderlyingStips,
			fix.NewKeyValue(FieldUnderlyingStipType, &fix.String{}),
			fix.NewKeyValue(FieldUnderlyingStipValue, &fix.String{}),
		),
	}
}

func (group *UnderlyingStipsGrp) AddEntry(entry *UnderlyingStipsEntry) *UnderlyingStipsGrp {
	group.Group.AddEntry(entry.Items())

	return group
}

func (group *UnderlyingStipsGrp) Entries() []*UnderlyingStipsEntry {
	items := make([]*UnderlyingStipsEntry, len(group.Group.Entries()))

	for i, item := range group.Group.Entries() {
		items[i] = &UnderlyingStipsEntry{fix.NewComponent(item...)}
	}

	return items
}

type UnderlyingStipsEntry struct {
	*fix.Component
}

func makeUnderlyingStipsEntry() *UnderlyingStipsEntry {
	return &UnderlyingStipsEntry{fix.NewComponent(
		fix.NewKeyValue(FieldUnderlyingStipType, &fix.String{}),
		fix.NewKeyValue(FieldUnderlyingStipValue, &fix.String{}),
	)}
}

func NewUnderlyingStipsEntry() *UnderlyingStipsEntry {
	return makeUnderlyingStipsEntry()
}

func (underlyingStipsEntry *UnderlyingStipsEntry) UnderlyingStipType() string {
	kv := underlyingStipsEntry.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (underlyingStipsEntry *UnderlyingStipsEntry) SetUnderlyingStipType(underlyingStipType string) *UnderlyingStipsEntry {
	kv := underlyingStipsEntry.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(underlyingStipType)
	return underlyingStipsEntry
}

func (underlyingStipsEntry *UnderlyingStipsEntry) UnderlyingStipValue() string {
	kv := underlyingStipsEntry.Get(1)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (underlyingStipsEntry *UnderlyingStipsEntry) SetUnderlyingStipValue(underlyingStipValue string) *UnderlyingStipsEntry {
	kv := underlyingStipsEntry.Get(1).(*fix.KeyValue)
	_ = kv.Load().Set(underlyingStipValue)
	return underlyingStipsEntry
}
