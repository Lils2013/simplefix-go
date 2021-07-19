package fix44

import (
	"github.com/b2broker/simplefix-go/fix"
)

type UnderlyingStipulations struct {
	*fix.Component
}

func makeUnderlyingStipulations() *UnderlyingStipulations {
	return &UnderlyingStipulations{fix.NewComponent(
		NewUnderlyingStipsGrp().Group,
	)}
}

func NewUnderlyingStipulations() *UnderlyingStipulations {
	return makeUnderlyingStipulations()
}

func (underlyingStipulations *UnderlyingStipulations) UnderlyingStipsGrp() *UnderlyingStipsGrp {
	group := underlyingStipulations.Get(0).(*fix.Group)

	return &UnderlyingStipsGrp{group}
}

func (underlyingStipulations *UnderlyingStipulations) SetUnderlyingStipsGrp(noUnderlyingStips *UnderlyingStipsGrp) *UnderlyingStipulations {
	underlyingStipulations.Set(0, noUnderlyingStips.Group)

	return underlyingStipulations
}