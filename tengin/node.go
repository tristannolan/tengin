package tengin

type Node interface {
	Canvas() *Canvas
	Control() *Control
}

type BaseNode struct {
	canvas  *Canvas
	control *Control
}

func (n BaseNode) Canvas() *Canvas {
	return n.canvas
}

func (n BaseNode) Control() *Control {
	return n.control
}

// ==================
//
//	Implementation
//
// ==================
type CustomNode struct {
	*BaseNode
	Label string
}

func NewCustomNode() *CustomNode {
	n := CustomNode{
		BaseNode: &BaseNode{},
		Label:    "My Custom Node",
	}

	return &n
}
