package gel

import (
	"bytes"
	"fmt"
	"io"
)

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

// Nodes represent the different parts of of Html as one type.  A single Node
// can be one and only one of the following types: Textual, Element, or
// Attribute.  In general a Textual node will not have children, but will
// have CData, Elements can have both children of type Text, Element,
// Attributes or Fragment, but cannot directly hold CData.  Attributes have
// only Key and Value strings, and all other fields are empty or nil.  Lastly,
// Fragments can have children of type Text and Element, while all other
// fields are empty or nil.
type Node struct {
	Tag      string
	Children []*Node
	Atts     []*Node
	Type     Type
	Key      string
	Value    string
	CData    string
	IsVoid   bool
}

// WriteTo will output the Node to the writer correctly nesting children and
// attributes.
func (e *Node) WriteTo(w io.Writer) {
	e.WriteToIndented(Indent{}, w)
}

// ToNode is implemented to conform to a component pattern of Nodes within
// Nodes, but additionally some other instance can pose as a View if they
// implement ToNode.
func (e *Node) ToNode() *Node {
	return e
}

// WriteTo writes the Node to the given writer with the given indention.
func (e *Node) WriteToIndented(in Indent, w io.Writer) {
	switch e.Type {
	case Textual:
		if in.HasIndent() && e.CData != "" {
			in.WriteTo(w)
		}
		if e.CData != "" {
			w.Write([]byte(e.CData))
		}
		if in.HasIndent() && e.CData != "" {
			w.Write([]byte("\n"))
		}
	case Attribute:
		w.Write([]byte(" "))
		w.Write([]byte(e.Key))
		w.Write([]byte("=\""))
		w.Write([]byte(e.Value))
		w.Write([]byte("\""))
	case Fragment:
		for _,f := range e.Children {
			f.WriteToIndented(in, w)
		}
	case Element:
		if in.HasIndent() {
			in.WriteTo(w)
		}
		w.Write([]byte("<"))
		w.Write([]byte(e.Tag))
		if len(e.Atts) > 0 {
			for _, att := range e.Atts {
				att.WriteTo(w)
			}
		}
		if e.IsVoid {
			w.Write([]byte("/>"))
		} else {
			w.Write([]byte(">"))
			if len(e.Children) > 0 {
				if in.HasIndent() {
					w.Write([]byte("\n"))
				}
				next := in.Incr()
				for _, kid := range e.Children {
					kid.WriteToIndented(next, w)
				}
			}
		}
		if in.Level > 0 && in.HasIndent() && len(e.Children) > 0 {
			in.WriteTo(w)
		}
		if !e.IsVoid {
			w.Write([]byte(fmt.Sprintf("</%s>", e.Tag)))
		}
		if in.Level > 0 {
			w.Write([]byte("\n"))
		}
	}
}

// String renders the Node as html (text, element, or attribute).
func (e *Node) String() string {
	buf := bytes.NewBuffer([]byte{})
	e.WriteTo(buf)
	return buf.String()
}

// Add will collect and bucket the nodes into atts and children.  Nodes
// of type Text or Element are added to children and Attribute type are
// added to the Atts slice.  If the Node is not an Element then
// attributes will silently be ignored.
func (t *Node) Add(nodes ...*Node) *Node {
	for _, n := range nodes {
		switch n.Type {
		case Textual, Element, Fragment:
			t.Children = append(t.Children, n)
		case Attribute:
			if t.Type == Element {
				t.Atts = append(t.Atts, n)
			}
		}
	}
	return t
}

// Text adds the given strings as text nodes.
func (n *Node) Text(ts ...string) *Node {
	for _, c := range ts {
		n.Add(Text(c))
	}
	return n
}

// Att creates a new Node with Attribute type and the given key, value pair.
func Att(key, value string) *Node {
	node := &Node{
		Type:  Attribute,
		Key:   key,
		Value: value,
	}
	return node
}

// Text creates a new Node with Textual type and CData from the provided string.
func Text(c string) *Node {
	node := &Node{
		Type:  Textual,
		CData: c,
	}
	return node
}

// Fmt creates a Text node using Sprintf.
func Fmt(format string, args ...interface{}) *Node {
	s := fmt.Sprintf(format, args...)
	return Text(s)
}

func Frag(children ...*Node) *Node {
	n := &Node{
		Type: Fragment,
		Children: make([]*Node, 0),
	}
	return n.Add(children...)
}

// None produces an empty Text Node.
func None() *Node {
	return Text("")
}

// Maybe check if val is nil, a view or a viewable, and if so returns
// a View based on the val, but if it's none of these then it returns
// an empty Text node.
func Maybe(val interface{}) View {
	if val == nil {
		return None()
	}
	viewable, ok := val.(Viewable)
	if ok {
		return viewable.ToView()
	}
	view, ok := val.(View)
	if ok {
		return view
	}
	return None()
}

