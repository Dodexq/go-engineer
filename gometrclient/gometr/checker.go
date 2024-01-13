package gometr

import (
	"fmt"
	"strings"
)

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

func (c *Checker) Add(elements ...Checkable) {
	c.items = append(c.items, elements...)
}

func (c *Checker) GetAllIDs() {
	for i := range c.items {
		fmt.Println(i)
	}
}

func (c *Checker) String() string {
	var ids []string
	for _, item := range c.items {
		ids = append(ids, item.GetID())
	}
	return fmt.Sprintf("Checker with IDs: %s", strings.Join(ids, ", "))
}

func (c *Checker) GetMetrics() string {
	return "This is a metrics"
}

func (c *Checker) Ping() error {
	return nil
}

func (c *Checker) GetID() string {
	var userInput string

	fmt.Print("Введите ID: ")
	fmt.Scanln(&userInput)

	return strings.TrimSpace(userInput)
}

func (c *Checker) Health() bool {
	return true
}
