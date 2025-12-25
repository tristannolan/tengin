package tengin

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/gdamore/tcell/v3"
)

var (
	debugMessages           = []debugMsg{}
	persistentDebugMessages = []debugMsg{}
	longestName             = 0
	longestValue            = 0
)

type debug struct {
	enabled bool
}

type debugMsg struct {
	name  string
	value string
}

func newDebug() debug {
	return debug{
		enabled: true,
	}
}

func newDebugMsg(name string, value any) debugMsg {
	msg := debugMsg{
		name: name,
	}

	switch v := value.(type) {
	case string:
		msg.value = v
	case int32:
		msg.value = strconv.Itoa(int(v))
	case int:
		msg.value = strconv.Itoa(v)
	case float32:
		msg.value = strconv.FormatFloat(float64(v), 'f', 2, 32)
	case float64:
		msg.value = strconv.FormatFloat(v, 'f', 2, 64)
	case bool:
		msg.value = strconv.FormatBool(v)
	default:
		msg.value = ""
	}

	if len(msg.name) > longestName {
		longestName = len(msg.name)
	}
	if len(msg.value) > longestValue {
		longestValue = len(msg.value)
	}

	return msg
}

func (d debug) update() {
	debugMessages = []debugMsg{}
}

func (d debug) draw(s tcell.Screen) {
	if d.enabled == false {
		return
	}

	msgs := slices.Concat(debugMessages, persistentDebugMessages)

	w, h := s.Size()
	x := w - longestName - longestValue - 1
	y := h - len(msgs)

	for i, msg := range msgs {
		whitespace := ""
		for range longestName - len(msg.name) {
			whitespace += " "
		}

		s.PutStr(x-1, y+i, fmt.Sprintf("%s%s %s", msg.name, whitespace, msg.value))
	}
}

func DebugLog(name string, value any) {
	msg := newDebugMsg(name, value)
	debugMessages = append(debugMessages, msg)
}

func PersistentDebugLog(name string, value any) {
	msg := newDebugMsg(name, value)
	persistentDebugMessages = append(persistentDebugMessages, msg)
}
