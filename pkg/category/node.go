package category

import "sort"

//Node is used to build a parent-child relationship amoung categories
type Node struct {
	category *Category
	parent   *Node
	children []*Node
	Priority int
}

//NewNode creates a new node holding the category value. It has methods for helping building the final tree
func NewNode(category *Category, priority int) *Node {
	return &Node{category: category, Priority: priority}
}

//Category returns the category hold by the node
func (c *Node) Category() *Category {
	return c.category
}

//Children returns the children nodes hold by the node
func (c *Node) Children() []*Node {
	return c.children
}

//HasChildren returns true if the length of children is > 0, otherwise returns false
func (c *Node) HasChildren() bool {
	return len(c.children) > 0
}

//Parent returns the parent node hold by the node
func (c *Node) Parent() *Node {
	return c.parent
}

//SetParent overrides the parent of a node, with another node
func (c *Node) SetParent(node *Node) {
	c.parent = node
}

//AddChild append a child node to the node
func (c *Node) AddChild(node *Node) {
	c.children = append(c.children, node)
}

//IsRoot checks if the node is root, a.k.a., has a nil parent
func (c *Node) IsRoot() bool {
	return c.parent == nil
}

//SortChildren sorts the children of the node based on the Priority value.
func (c *Node) SortChildren() {
	sort.Slice(c.children, func(i, j int) bool { return c.children[i].Priority < c.children[j].Priority })
}
