package tengin

import (
	"cmp"
	"slices"
)

type controlManager struct {
	controls []*Control
	dirty    bool
}

type Control struct {
	width, height int
	x, y, z       int
	bounds        Rect
	dirty         bool
	Action        func()
	manager       *controlManager
}

func newControlManager() *controlManager {
	return &controlManager{
		controls: []*Control{},
		dirty:    true,
	}
}

func NewControl(x, y, width, height int, action func()) *Control {
	return &Control{
		x:       x,
		y:       y,
		Action:  action,
		dirty:   true,
		bounds:  NewRect(x, y, x+width-1, y+height-1),
		manager: nil,
	}
}

func (cm *controlManager) AppendControl(c ...*Control) {
	for _, ctrl := range c {
		ctrl.assignManager(cm)
		cm.controls = append(cm.controls, ctrl)
	}
}

func (cm *controlManager) RemoveControl(c ...*Control) {
	if len(cm.controls) == 0 {
		return
	}

	toRemove := make(map[*Control]struct{}, len(c))
	toRemain := make([]*Control, len(cm.controls)-len(c))

	for _, control := range c {
		toRemove[control] = struct{}{}
	}

	for _, control := range cm.controls {
		if _, found := toRemove[control]; found {
			continue
		}
		toRemain = append(toRemain, control)
	}
}

func (cm controlManager) IsDirty() bool {
	return cm.dirty
}

func (cm *controlManager) Sort() {
	slices.SortStableFunc(cm.controls, func(a, b *Control) int {
		return cmp.Compare(a.z, b.z)
	})
}

func (c *Control) assignManager(cm *controlManager) {
	c.manager = cm
}

func (c Control) Z() int {
	return c.z
}

func (c *Control) SetZ(z int) {
	c.z = z
	c.markDirty()
}

func (c *Control) IsDirty() bool {
	return c.dirty
}

func (c *Control) markDirty() {
	if c.dirty == true {
		return
	}

	c.dirty = true
	if c.manager.IsDirty() {
		c.manager.dirty = true
	}
}

func (c *Control) ContainsPoint(x, y int) bool {
	return c.bounds.Contains(x, y)
}
