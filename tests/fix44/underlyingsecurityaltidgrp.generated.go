package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
)

type UnderlyingSecurityAltIDGrp struct {
	*fix.Group
}

func NewUnderlyingSecurityAltIDGrp() *UnderlyingSecurityAltIDGrp {
	return &UnderlyingSecurityAltIDGrp{
		fix.NewGroup(FieldNoUnderlyingSecurityAltID,
			fix.NewKeyValue(FieldUnderlyingSecurityAltID, &fix.String{}),
			fix.NewKeyValue(FieldUnderlyingSecurityAltIDSource, &fix.String{}),
		),
	}
}

func (group *UnderlyingSecurityAltIDGrp) AddEntry(entry *UnderlyingSecurityAltIDEntry) *UnderlyingSecurityAltIDGrp {
	group.Group.AddEntry(entry.Items())

	return group
}

func (group *UnderlyingSecurityAltIDGrp) Entries() []*UnderlyingSecurityAltIDEntry {
	items := make([]*UnderlyingSecurityAltIDEntry, len(group.Group.Entries()))

	for i, item := range group.Group.Entries() {
		items[i] = &UnderlyingSecurityAltIDEntry{fix.NewComponent(item...)}
	}

	return items
}

type UnderlyingSecurityAltIDEntry struct {
	*fix.Component
}

func makeUnderlyingSecurityAltIDEntry() *UnderlyingSecurityAltIDEntry {
	return &UnderlyingSecurityAltIDEntry{fix.NewComponent(
		fix.NewKeyValue(FieldUnderlyingSecurityAltID, &fix.String{}),
		fix.NewKeyValue(FieldUnderlyingSecurityAltIDSource, &fix.String{}),
	)}
}

func NewUnderlyingSecurityAltIDEntry() *UnderlyingSecurityAltIDEntry {
	return makeUnderlyingSecurityAltIDEntry()
}

func (underlyingSecurityAltIDEntry *UnderlyingSecurityAltIDEntry) UnderlyingSecurityAltID() string {
	kv := underlyingSecurityAltIDEntry.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (underlyingSecurityAltIDEntry *UnderlyingSecurityAltIDEntry) SetUnderlyingSecurityAltID(underlyingSecurityAltID string) *UnderlyingSecurityAltIDEntry {
	kv := underlyingSecurityAltIDEntry.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(underlyingSecurityAltID)
	return underlyingSecurityAltIDEntry
}

func (underlyingSecurityAltIDEntry *UnderlyingSecurityAltIDEntry) UnderlyingSecurityAltIDSource() string {
	kv := underlyingSecurityAltIDEntry.Get(1)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (underlyingSecurityAltIDEntry *UnderlyingSecurityAltIDEntry) SetUnderlyingSecurityAltIDSource(underlyingSecurityAltIDSource string) *UnderlyingSecurityAltIDEntry {
	kv := underlyingSecurityAltIDEntry.Get(1).(*fix.KeyValue)
	_ = kv.Load().Set(underlyingSecurityAltIDSource)
	return underlyingSecurityAltIDEntry
}
