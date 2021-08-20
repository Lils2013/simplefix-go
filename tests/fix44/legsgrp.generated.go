package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
)

type LegsGrp struct {
	*fix.Group
}

func NewLegsGrp() *LegsGrp {
	return &LegsGrp{
		fix.NewGroup(FieldNoLegs,
			makeInstrumentLeg().Component,
		),
	}
}

func (group *LegsGrp) AddEntry(entry *LegsEntry) *LegsGrp {
	group.Group.AddEntry(entry.Items())

	return group
}

func (group *LegsGrp) Entries() []*LegsEntry {
	items := make([]*LegsEntry, len(group.Group.Entries()))

	for i, item := range group.Group.Entries() {
		items[i] = &LegsEntry{fix.NewComponent(item...)}
	}

	return items
}

type LegsEntry struct {
	*fix.Component
}

func makeLegsEntry() *LegsEntry {
	return &LegsEntry{fix.NewComponent(
		makeInstrumentLeg().Component,
	)}
}

func NewLegsEntry() *LegsEntry {
	return makeLegsEntry()
}

func (legsEntry *LegsEntry) InstrumentLeg() *InstrumentLeg {
	component := legsEntry.Get(0).(*fix.Component)

	return &InstrumentLeg{component}
}

func (legsEntry *LegsEntry) SetInstrumentLeg(instrumentLeg *InstrumentLeg) *LegsEntry {
	legsEntry.Set(0, instrumentLeg.Component)

	return legsEntry
}
