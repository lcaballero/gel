package gel

// Fragment is the concrete type of the Views interface.
type Fragment []View

// Frag creates a view that is a list of children views without a
// containing element.
func Frag(children ...View) Fragment {
	return children
}

// NewViews allocates an empty views list to which move views can be
// added.
func NewFragment() Fragment {
	return Fragment{}
}

// Add appends the given views to the currently established list of
// views.
func (v Fragment) Add(views ...View) Fragment {
	return append(v, views...)
}

// ToViews implements to the Viewable interface returning a concrete
// Node fragment instance.
func (v Fragment) ToNode() *Node {
	n := &Node{
		Type:     NodeList,
		Children: []*Node{},
	}
	return n.Add(v...).ToNode()
}

// Len returns the number of views held by the Fragment.
func (v Fragment) Len() int {
	return len(v)
}
