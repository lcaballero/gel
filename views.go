package gel

// View interface represents anything that can be turned into a Node via the
// ToNode() function.
type View interface {
	ToNode() *Node
}
