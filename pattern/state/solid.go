package main

import "fmt"

type SolidState struct {
	name string
}

func NewSolidState() *SolidState {
	return &SolidState{name: "solid"}
}

func (s *SolidState) getName() string {
	return s.name
}

func (s *SolidState) freeze(c *StateContext) {
	fmt.Println("Nothing happens.")
}

func (s *SolidState) heat(c *StateContext) {
	c.setState(NewLiquidState())
}
