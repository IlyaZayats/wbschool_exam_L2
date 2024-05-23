package main

import "fmt"

type Strategy interface {
	use()
}

type StrategyA struct{}

func (s *StrategyA) use() {
	fmt.Println("StrategyA")
}

type StrategyB struct{}

func (s *StrategyB) use() {
	fmt.Println("StrategyB")
}

type StrategyC struct{}

func (s *StrategyC) use() {
	fmt.Println("StrategyC")
}

type Context struct {
	strategy Strategy
}

func (c *Context) useStrategy() {
	c.strategy.use()
}

func (c *Context) setStrategy(newStrategy Strategy) {
	c.strategy = newStrategy
}

func main() {
	context := Context{}
	strategies := []Strategy{&StrategyA{}, &StrategyB{}, &StrategyC{}}
	for _, strategy := range strategies {
		context.setStrategy(strategy)
		context.useStrategy()
	}
}
