package tengin

// everything here is currently a placeholder, general implementation only

// Node holds entities and their functions.
// Node function: layout - dimensions, transform
// Node function: hold children and parent for positionally relative nesting
// Node function: canvas - draw an image (optional)
// Node function: control - react to input (optional)

// Nodes are submitted to a scene, where they are processed and rendered.
// Nodes with a canavs will be rendered to the screen (if applicable)
// Nodes with a control will react to input (if applicable)
// The scene must be active for nodes to have an effect. They can be stored in
// an inactive scene to preload or hotswap.

// A scene is purely a storage device for nodes, and a functional renderer
// ECS style management must come externally
// The main game struct should contain a world, which submits an active scene
// World will coordinate updates by looping through nodes.
// E.g. world => node.Canvas().Update()
// E.g. world => node.Control().HitTest(mouseX, mouseY) [how to trigger universally]
type Node struct {
	ID         int
	components []Component
}

type Component interface {
	Node() *Node
}

type CanvasComponent struct {
	node   *Node
	canvas *Canvas
}

type ControlComponent struct {
	node    *Node
	control *Control
}

func (cc CanvasComponent) Node() *Node {
	return cc.node
}

func (cc ControlComponent) Node() *Node {
	return cc.node
}

type ButtonComponent struct {
	node *Node
}

func NewButtonComponent() Node {
	node := Node{
		ID: 1,
		components: []Component{
			CanvasComponent{},
			ControlComponent{},
		},
	}

	return node
}
