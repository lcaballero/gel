package gel

// Views represent a list of views that can be rendered as a one View.
type Views []View

// NewViews allocates an empty views list to which move views can be added.
func NewViews() Views {
	return make([]View, 0)
}

// Add appends the given views to the currently established list of views.
func (v Views) Add(views ...View) Views {
	v = append(v, views...)
	return v
}

// ToViews implements to the Viewable interface returning a concrete Node
// fragment instance.
func (v Views) ToView() View {
	return Frag(v...)
}

// ToNode implements the View interface, producing the final transformation
// to a *Node which can be rendered to stream.
func (v Views) ToNode() *Node {
	return v.ToView().ToNode()
}

// Viewable is the interface for anything that can be rendered into a View.
type Viewable interface {
	ToView() View
}

// ToView is the functional equivalent of Viewable, which allows an
// empty function returning a View to be a Viewable.
type ToView func() View

// ToView method implments the Viewable interface for the ToView func.
func (v ToView) ToView() View {
	return v()
}

// View interface represents anything that can be turned into a Node via the
// ToNode() function.
type View interface {
	ToNode() *Node
}

// ToNode is an empty function that returns a rendered/completed Node.
type ToNode func() *Node

// ToNode here implements the View interface, meaning that a ToNode function
// can now be used like a View.
func (t ToNode) ToNode() *Node {
	return t()
}

