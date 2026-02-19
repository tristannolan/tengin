# Plans
## Interface away
More elements should work through interfaces. Notably, the "role" functions.

type Updatable inteface
type Drawable interface
type ClickHandler interface
type Transformable interface
type Dirtyable interface

This allows me to make multiple things work the same way. For example, a node,
canvas, and control can all get the update and draw functions just by including
the update and draw interfaces.

A key question is "could this element function in different ways in the future?"
If there's no chance, or the element is data only, then it's a struct. If it's
likely that a different module or script could submit a result, then an
interface might not be a bad choice.
