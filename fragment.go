package gel

// Fragment is the concrete type of the Views interface.
type Fragment []View

// Add appends the given views to the currently established list of
// views.
func (v Fragment) Add(views ...View) Fragment {
	return append(v, views...)
}

// ToViews implements to the Viewable interface returning a concrete
// Node fragment instance.
func (v Fragment) ToNode() *Node {
	n := &Node{
		kind: NodeList,
	}
	return n.Add(v...).ToNode()
}

// Len returns the number of views held by the Fragment.
func (v Fragment) Len() int {
	return len(v)
}
