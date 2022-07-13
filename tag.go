package gel

// E creates an element that can hold children.
func E(tag string) Tag {
	return El(tag, false)
}

// El creates a Tag func for the given tag name and isVoid flag.
func El(tag string, isVoid bool) Tag {
	return func(children ...View) View {
		node := &Node{
			Type:       Element,
			Tag:        tag,
			Children:   make([]*Node, 0),
			Attributes: make([]*Node, 0),
			IsVoid:     isVoid,
		}
		node.Add(children...)
		return node
	}
}

// Tag is the starting point of an element.
type Tag func(...View) View

// ToNode renders the Tag as a Node.
func (t Tag) ToNode() *Node {
	return t().ToNode()
}

// Adds the class attribute with the given value
func (t Tag) Class(class string) Tag {
	return t.Atts("class", class)
}

// Atts creates a tag with the given pairs of Attributes.
func (t Tag) Atts(pairs ...string) Tag {
	return func(views ...View) View {
		atts := []View{Atts(pairs...)}
		atts = append(atts, views...)
		return t(atts...)
	}
}

// Text will create an Element Node from the Tag and then immediately add the
// given strings as Text nodes.
func (t Tag) Text(c ...string) View {
	e := t().ToNode()
	for _, txt := range c {
		e.Add(Text(txt))
	}
	return e
}

// Add appends the children provided as views to the current tag.
func (t Tag) Add(children ...View) View {
	return t().ToNode().Add(children...)
}

// Fmt creates a Text tag using the given format and args like Sprintf.
func (t Tag) Fmt(format string, args ...interface{}) View {
	return t(Fmt(format, args...))
}
