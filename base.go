package gel

type Tag func(...*Node) *Node

// el creates a Tag func for the given tag name and isVoid flag.
func el(tag string, isVoid bool) Tag {
	return func(children ...*Node) *Node {
		node := &Node{
			Type: Element,
			Tag:  tag,
			Children: make([]*Node, 0),
			Atts: make([]*Node, 0),
			IsVoid: isVoid,
	}
		node.Add(children...)
		return node
	}
}

// Text will create an Element Node from the Tag and then immediately add the
// given strings as Text nodes.
func (t Tag) Text(c ...string) *Node {
	e := t()
	for _,txt := range c {
		e.Add(Text(txt))
	}
	return e
}

// Add allocates a Node with the given Tag name (lower-cased) and the provided
// children.  This is an alias to tag.New(...).
func (t Tag) Add(children ...*Node) *Node {
	return t(children...)
}





