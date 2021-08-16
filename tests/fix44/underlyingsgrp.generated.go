package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
)

type UnderlyingsGrp struct {
	*fix.Group
}

func NewUnderlyingsGrp() *UnderlyingsGrp {
	return &UnderlyingsGrp{
		fix.NewGroup(FieldNoUnderlyings,
			makeUnderlyingInstrument().Component,
		),
	}
}

func (group *UnderlyingsGrp) AddEntry(entry *UnderlyingsEntry) *UnderlyingsGrp {
	group.Group.AddEntry(entry.Items())

	return group
}

func (group *UnderlyingsGrp) Entries() []*UnderlyingsEntry {
	items := make([]*UnderlyingsEntry, len(group.Group.Entries()))

	for i, item := range group.Group.Entries() {
		items[i] = &UnderlyingsEntry{fix.NewComponent(item...)}
	}

	return items
}

type UnderlyingsEntry struct {
	*fix.Component
}

func makeUnderlyingsEntry() *UnderlyingsEntry {
	return &UnderlyingsEntry{fix.NewComponent(
		makeUnderlyingInstrument().Component,
	)}
}

func NewUnderlyingsEntry() *UnderlyingsEntry {
	return makeUnderlyingsEntry()
}

func (underlyingsEntry *UnderlyingsEntry) UnderlyingInstrument() *UnderlyingInstrument {
	component := underlyingsEntry.Get(0).(*fix.Component)

	return &UnderlyingInstrument{component}
}

func (underlyingsEntry *UnderlyingsEntry) SetUnderlyingInstrument(underlyingInstrument *UnderlyingInstrument) *UnderlyingsEntry {
	underlyingsEntry.Set(0, underlyingInstrument.Component)

	return underlyingsEntry
}
