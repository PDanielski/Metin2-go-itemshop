package category

//NodeTreeProvider allows a polymorphic approach for retrieving node trees.
type NodeTreeProvider interface {
	Provide() (*NodeTree, error)
}

//NodeTree is a collection of root nodes
type NodeTree struct {
	rootNodes []*Node
}

//Roots returns the collection of root nodes
func (nt *NodeTree) Roots() []*Node {
	return nt.rootNodes
}

//NewNodeTree constructs a new node tree with the root nodes given in the input
func NewNodeTree(nodes []*Node) *NodeTree {
	return &NodeTree{nodes}
}
