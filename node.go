package gel

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type View interface {
	ToNode() *Node
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
	Tag      Tag
	Children []*Node
	Atts     []*Node
	Type     Type
	Key      string
	Value    string
	CData    string
}

// WriteTo will output the Node to the writer correctly nesting children and
// attributes.
func (e *Node) WriteTo(w io.Writer) {
	e.WriteToIndented(Indent{Level: 0}, w)
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
		tag := strings.ToLower(e.Tag.String())
		w.Write([]byte(tag))
		if len(e.Atts) > 0 {
			for _, att := range e.Atts {
				att.WriteTo(w)
			}
		}
		if e.Tag.IsSelfClosing() {
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
		if !e.Tag.IsSelfClosing() {
			w.Write([]byte(fmt.Sprintf("</%s>", tag)))
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

func Frag(children ...*Node) *Node {
	n := &Node{
		Type: Fragment,
		Children: make([]*Node, 0),
	}
	return n.Add(children...)
}

func None() *Node {
	return Text("")

}

// New allocates a Node with the provided 0 or more children and the lower-case
// name of the Tag.
func (t Tag) New(children ...*Node) *Node {
	node := &Node{
		Type: Element,
		Tag:  t,
		Children: make([]*Node, 0),
		Atts: make([]*Node, 0),
	}
	node.Add(children...)
	return node
}

// Text will create an Element Node from the Tag and then immediately add the
// given strings as Text nodes.
func (t Tag) Text(c ...string) *Node {
	e := t.New()
	for _,txt := range c {
		e.Add(Text(txt))
	}
	return e
}

// Add allocates a Node with the given Tag name (lower-cased) and the provided
// children.  This is an alias to tag.New(...).
func (t Tag) Add(children ...*Node) *Node {
	return t.New(children...)
}



