package tengin

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v3"
)

var (
	debugMessages = []debugMsg{}
	longestName   = 0
	longestValue  = 0
)

type debug struct {
	enabled bool
}

func newDebug() debug {
	return debug{
		enabled: true,
	}
}

type debugMsg struct {
	name  string
	value string
}

func DebugLog(name string, value any) {
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
	case bool:
		msg.value = strconv.FormatBool(v)
	}

	if len(msg.name) > longestName {
		longestName = len(msg.name)
	}
	if len(msg.value) > longestValue {
		longestValue = len(msg.value)
	}

	debugMessages = append(debugMessages, msg)
}

func (d debug) update() {
	debugMessages = []debugMsg{}
}

func (d debug) draw(s tcell.Screen) {
	if d.enabled == false {
		return
	}

	w, h := s.Size()
	x := w - longestName - longestValue - 1
	y := h - len(debugMessages)

	for i, msg := range debugMessages {
		whitespace := ""
		for range longestName - len(msg.name) {
			whitespace += " "
		}

		s.PutStr(x-1, y+i, fmt.Sprintf("%s%s %s", msg.name, whitespace, msg.value))
	}
}
