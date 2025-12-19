package tengin

import "github.com/gdamore/tcell/v3"

// A tile holds the actual graphical information
// It doesn't hold its positional information because it only delivers
// the rendering information, the rest can be handled in storage systems
type Tile struct {
	char       string
	background tcell.Style
	text       tcell.Style
}

// A canvas holds graphical information. It's effectively a box of tiles with a
// defined width and location
// It could be an image, a textbox, a paragraph, group of text boxes, each holding
// their own relative canvas text boxes. On that note, Canvases should actually be
// able to hold their own children. It might be preferable when creating a ui
// or something, because we'd want to group textbox canvases together into a menu
// It would be a pain to put those all directly onto scene and manage them individually.
// Then I could deliver canvas to scene once when I'm happy that I've assembled the
// entire thing. When scene renders it can compile the children, progressively
// moving higher until it's just connecting canvases. Then it renders
// We could add text, or make it a background colour, or draw a picture
// Translation functions give animation power
// A z axis allows for layered utility in rendering order, the engine can optimise
// which tiles get drawn first when a collection of tiles are delivered for render
// I don't like the name though, the plural doesn't sound nice. Image? Meh
type Canvas struct {
	x, y, z       int
	width, height int
	tiles         []Tile
	children      []Canvas
}

func (c *Canvas) Rotate() {
}

// Canvas primatives
// these would add some input processing to produce common canvases
type (
	Textbox   Canvas
	Paragraph Canvas
	Columns   Canvas
	Box       Canvas
	Gradient  Canvas
)

// A scene is a special type of canvas. It works the same way in that it can hold
// canvas structs with positions relative to it's own. It has the added benefit of
// compiling the internal canvases together. Passing the scene to the e.RenderScene()
// will let the engine perform tile diffing and other fancy stuff to optimise
// the render process
type Scene struct {
	x, y, z       int
	width, height int
	canvases      []Canvas
}

//                                                                                                                                         Input: f  Tick: 19820

// This might be a better solution
// ctx.RenderScene(s *Scene)
// Only one scene can be delivered to the render function. To perform something
// like split screen, two canvases should be made since they are effectively the
// same thing. Scene requires all information to be present for diffing
func (s *Scene) Render(scr tcell.Screen) {
	// 1. Collect visible tiles
	// 2. Resolve z-order
	// 3. Translate styles
	// 4. Batch optimally
	// 5. Write to screen
}
