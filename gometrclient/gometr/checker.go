package gometr

import "fmt"

type Checker struct {
	items []Checkable
}

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

type Metrics struct {
	ID string
}

func (c *Checker) Add(elements ...Checkable) {
	c.items = append(c.items, elements...)
}

func (c *Checker) GetAllIDs() {
	for i := range c.items {
		fmt.Println(i)
	}
}

func (m *Metrics) GetMetrics() string {
	return "This is a metrics"
}

func (m *Metrics) Ping() error {
	return nil
}

func (m *Metrics) GetID() string {
	return m.ID
}

func (m *Metrics) Health() bool {
	return true
}
