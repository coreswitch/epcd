package pgw

import (
	"github.com/coreswitch/component"
)

// Component is PGW component.
type Component struct {
}

// Start is component's start function.
func (c *Component) Start() component.Component {
	return c
}

// Stop is component's stop function.
func (c *Component) Stop() component.Component {
	return c
}
