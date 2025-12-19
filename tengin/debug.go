package tengin

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v3"
)

type debug struct {
	enabled      bool
	messages     []debugMsg
	longestName  int
	longestValue int
}

func newDebug() debug {
	return debug{
		enabled:  true,
		messages: []debugMsg{},
	}
}

type debugMsg struct {
	name  string
	value string
}

func (d *debug) log(name string, value any) {
	msg := debugMsg{
		name: name,
	}

	switch v := value.(type) {
	case string:
		msg.value = v
	case int:
		msg.value = strconv.Itoa(v)
	case float32:
		msg.value = strconv.FormatFloat(float64(v), 'f', 2, 32)
	case float64:
		msg.value = strconv.FormatFloat(v, 'f', 2, 64)
	}

	if len(msg.name) > d.longestName {
		d.longestName = len(msg.name)
	}
	if len(msg.value) > d.longestValue {
		d.longestValue = len(msg.value)
	}

	d.messages = append(d.messages, msg)
}

func (d *debug) update() {
	d.messages = []debugMsg{}
}

func (d *debug) draw(s tcell.Screen) {
	if d.enabled == false {
		return
	}

	w, h := s.Size()
	x := w - d.longestName - d.longestValue - 1
	y := h - len(d.messages)

	for i, msg := range d.messages {
		whitespace := ""
		for range d.longestName - len(msg.name) {
			whitespace += " "
		}

		s.PutStr(x-1, y+i, fmt.Sprintf("%s%s %s", msg.name, whitespace, msg.value))
	}
}
