package main

type LiquidState struct {
	name string
}

func NewLiquidState() *LiquidState {
	return &LiquidState{name: "liquid"}
}

func (s *LiquidState) getName() string {
	return s.name
}

func (s *LiquidState) freeze(c *StateContext) {
	c.setState(NewSolidState())
}

func (s *LiquidState) heat(c *StateContext) {
	c.setState(NewGaseousState())
}
