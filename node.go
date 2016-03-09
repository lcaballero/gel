package gel

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Nodes represent the different parts of of Html as one type.  A single Node
// can be one and only one of the following types: Textual, Element, or
// Attribute.  In general a Textual node will not have children, but will
// have CData, Elements can have both children either of type Text or Element
// and Attributes, but does not directly hold CData.  Attributes have only
// Key and Value strings, and all other fields are empty or nil.
type Node struct {
	Tag      string
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

// WriteTo writes the Node to the given writer with the given indention.
func (e *Node) WriteToIndented(in Indent, w io.Writer) {
	switch e.Type {
	case Textual:
		w.Write([]byte(e.CData))
	case Attribute:
		w.Write([]byte(" "))
		w.Write([]byte(e.Key))
		w.Write([]byte("=\""))
		w.Write([]byte(e.Value))
		w.Write([]byte("\""))
	case Element:
		w.Write([]byte("<"))
		w.Write([]byte(e.Tag))
		if len(e.Atts) > 0 {
			for _, att := range e.Atts {
				att.WriteTo(w)
			}
		}
		w.Write([]byte(">"))
		if len(e.Children) > 0 {
			for _, kid := range e.Children {
				kid.WriteTo(w)
			}
		}
		w.Write([]byte(fmt.Sprintf("</%s>", e.Tag)))
	}
}

func (e *Node) String() string {
	buf := bytes.NewBuffer([]byte{})
	e.WriteTo(buf)
	return buf.String()
}

func (t *Node) Add(children ...*Node) *Node {
	for _, n := range children {
		switch n.Type {
		case Textual:
			t.Children = append(t.Children, n)
		case Element:
			t.Children = append(t.Children, n)
		case Attribute:
			t.Atts = append(t.Atts, n)
		}
	}
	return t
}

func Att(key, value string) *Node {
	node := &Node{
		Type:  Attribute,
		Key:   key,
		Value: value,
	}
	return node
}

func Text(c string) *Node {
	node := &Node{
		Type:  Textual,
		CData: c,
	}
	return node
}

func (t Tag) New(children ...*Node) *Node {
	node := &Node{
		Type: Element,
		Tag:  strings.ToLower(t.String()),
	}
	node.Add(children...)
	return node
}

func (t Tag) Add(children ...*Node) *Node {
	return t.New().Add(children...)
}
