package gel

type Tag func(...View) View

// el creates a Tag func for the given tag name and isVoid flag.
func el(tag string, isVoid bool) Tag {
	return func(children ...View) View {
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
func (t Tag) Text(c ...string) View {
	e := t().ToNode()
	for _,txt := range c {
		e.Add(Text(txt))
	}
	return e
}

// Fmt creates a Text tag using the given format and args like Sprintf.
func (t Tag) Fmt(format string, args ...interface{}) View {
	return t(Fmt(format, args...))
}
