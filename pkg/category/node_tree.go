package category

//NodeTree implementors allows to retrieve all the top level root categories
type NodeTree struct {
	rootNodes []*Node
}

//Roots return category node tree roots
func (nt *NodeTree) Roots() []*Node {
	return nt.rootNodes
}

//NodeTreeProvider implementation are used to retrieve NodeTrees
type NodeTreeProvider interface {
	Provide() (*NodeTree, error)
}

//NewNodeTree constructs a new node tree given root nodes
func NewNodeTree(nodes []*Node) *NodeTree {
	return &NodeTree{nodes}
}
