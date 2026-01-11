package cmd

import (
	"fmt"
	"strings"
)

type Controller struct {
	valid   map[Name]*Command
	buffer  string
	Trigger string
}

type Command struct {
	Phrase Name
	Action func([]string)
}

type Name string

const (
	cmdSet Name = "set"
)

func NewController() *Controller {
	return &Controller{
		valid:   map[Name]*Command{},
		buffer:  "",
		Trigger: ":",
	}
}

func New(phrase string, action func(args []string)) *Command {
	return &Command{
		Phrase: Name(phrase),
		Action: action,
	}
}

func (c Controller) Buffer() string {
	return c.buffer
}

func (c *Controller) ClearBuffer() {
	c.buffer = c.buffer[:0]
}

func (c *Controller) AppendToBuffer(part string) {
	c.buffer += part
}

func (c *Controller) RemoveFromBuffer(count int) {
	if len(c.buffer)-count <= 0 {
		c.ClearBuffer()
		return
	}
	c.buffer = c.buffer[:len(c.buffer)-count]
}

func (c *Controller) Execute() {
	phrases := strings.Split(c.buffer, " ")

	for _, phrase := range phrases {
		cmd, err := c.Match(Name(phrase))
		if err != nil {
			continue
		}
		cmd.Action(phrases)
	}
	c.ClearBuffer()
}

func (c *Controller) Register(commands ...*Command) error {
	for _, cmd := range commands {
		if _, err := c.Match(cmd.Phrase); err == nil {
			return fmt.Errorf("Phrase already exists: %s", cmd.Phrase)
		}
		c.valid[cmd.Phrase] = cmd
	}
	return nil
}

func (c *Controller) Match(phrase Name) (*Command, error) {
	if _, ok := c.valid[phrase]; ok {
		return c.valid[phrase], nil
	}
	return nil, fmt.Errorf("Phrase unknown: %s", phrase)
}
