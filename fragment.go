package gel

// Fragment is the concrete type of the Views interface.
type Fragment struct {
	views []View
}

// Frag creates a view that is a list of children views without a containing
// element.
func Frag(children ...View) View {
	n := &Node{
		Type:     NodeList,
		Children: make([]*Node, 0),
	}
	return n.Add(children...)
}

// NewViews allocates an empty views list to which move views can be added.
func NewFragment() *Fragment {
	return &Fragment{
		views: make([]View, 0),
	}
}

// Add appends the given views to the currently established list of views.
func (v *Fragment) Add(views ...View) View {
	v.views = append(v.views, views...)
	return v
}

// ToViews implements to the Viewable interface returning a concrete Node
// fragment instance.
func (v *Fragment) ToView() View {
	return Frag(v.views...)
}

// ToNode implements the View interface, producing the final transformation
// to a *Node which can be rendered to stream.
func (v Fragment) ToNode() *Node {
	return v.ToView().ToNode()
}

// Len returns the number of views held by the Fragment.
func (v Fragment) Len() int {
	return len(v.views)
}
