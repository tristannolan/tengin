package tengin

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v3"
)

var (
	debugMessages           = []debugMsg{}
	persistentDebugMessages = []debugMsg{}
	debugTimers             = map[int]*debugTimer{}
	longestName             = 0
	longestValue            = 0
	nextDebugTimerId        = 0
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

func (d debug) draw(s tcell.Screen) {
	if d.enabled == false {
		return
	}

	for i := range nextDebugTimerId {
		if _, ok := debugTimers[i]; !ok {
			continue
		}

		t := debugTimers[i]
		DebugLog(t.name, t.GetMsg())
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

	debugMessages = []debugMsg{}
}

func DebugLog(name string, value any) {
	msg := newDebugMsg(name, value)
	debugMessages = append(debugMessages, msg)
}

func PersistentDebugLog(name string, value any) {
	msg := newDebugMsg(name, value)
	persistentDebugMessages = append(persistentDebugMessages, msg)
}

type debugTimer struct {
	name             string
	id               int
	maxLogs          int
	logCount         int
	total, lastTotal time.Duration
	start, end       time.Time
}

func NewDebugTimer(name string) *debugTimer {
	t := &debugTimer{
		id:        nextDebugTimerId,
		name:      name,
		maxLogs:   20,
		logCount:  0,
		total:     0,
		lastTotal: 0,
	}
	debugTimers[t.id] = t
	nextDebugTimerId++
	return t
}

func (t *debugTimer) Start() {
	t.start = time.Now()
}

func (t *debugTimer) End() {
	t.end = time.Now()

	t.total += t.end.Sub(t.start)

	if t.logCount <= t.maxLogs {
		t.logCount++
		return
	}

	t.lastTotal = t.total
	t.total = 0
	t.logCount = 0
}

func (t *debugTimer) GetMsg() string {
	return fmt.Sprintf("%s", t.lastTotal/time.Duration(t.maxLogs))
}
