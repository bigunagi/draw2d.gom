package layout

import ()

// A tree have n root
// Each node have children represented by the firstChild followed by nextSiblings
// see Encoding general trees as binary trees in http://en.wikipedia.org/wiki/Binary_tree
// Each node have a reference to an Element
type Node struct {
	element  	Element
	parent      *Node
	firstChild  *Node
	nextSibling *Node
}

func NewNode(element Element) *Node {
	return &Node{element, nil, nil, nil}
}

func (node *Node) Element() Element {
	return node.element
}

func (node *Node) Parent() *Node {
	return node.parent
}

func (node *Node) SetChildrenElements(elements ...Element) {
	if node.firstChild != nil {
		node.firstChild.Detach()
	}
	if len(elements) > 0 {
		child := NewNode(elements[0])
		node.firstChild = child
		child.parent = node
		for _, element := range elements[1:] {
			nextChild := NewNode(element)
			child.nextSibling = nextChild
			nextChild.parent = node
			child = nextChild
		}
	}
}

func (node *Node) FirstChild() *Node {
	return node.firstChild
}

func (node *Node) LastChild() *Node {
	if node.firstChild != nil {
		// cursor is the first child of node
		cursor := node.firstChild
		for cursor.nextSibling != nil {
			cursor = cursor.nextSibling
		}
		// now cursor is the last child of node
		return cursor
	}
	return nil
}

func (node *Node) PreviousSibling() *Node {
	if node.parent != nil {
		cursor := node.parent.firstChild
		if cursor == node {
			// special case
			return nil
		} else {
			for cursor.nextSibling != node {
				cursor = cursor.nextSibling
			}
			// cursor is now the previous sibling of node
			return cursor
		}
	}
	return nil
}

func (node *Node) NextSibling() *Node {
	return node.nextSibling
}

func (node *Node) Root() *Node {
	root := node
	for root.parent != nil {
		root = root.parent
	}
	return root
}

func (node *Node) Depth() int {
	cursor := node.parent
	depth := 0
	for cursor != nil {
		depth++
		cursor = cursor.parent
	}
	return 0
}

func (node *Node) Detach() {
	if node.parent != nil {
		cursor := node.parent.firstChild
		if cursor == node {
			// special case
			node.parent.firstChild = node.nextSibling
		} else {
			for cursor.nextSibling != node {
				cursor = cursor.nextSibling
			}
			// cursor is now the previous sibling of node
			cursor.nextSibling = node.nextSibling
		}
		node.nextSibling = nil
		node.parent = nil
	}
}

func (refChild *Node) InsertBefore(newChild *Node) {
	// can't insert before if refChild have no parent
	if refChild.parent != nil {
		newChild.Detach()
		parent := refChild.parent
		cursor := parent.firstChild
		if cursor == refChild {
			// special case
			parent.firstChild = newChild
		} else {
			for cursor.nextSibling != refChild {
				cursor = cursor.nextSibling
			}
			// cursor is now the previous sibling of node
			cursor.nextSibling = newChild
		}
		newChild.parent = parent
		newChild.nextSibling = refChild
	}
}

func (refChild *Node) InsertAfter(newChild *Node) {
	newChild.Detach()
	newChild.parent = refChild.parent
	newChild.nextSibling = refChild.nextSibling
	refChild.nextSibling = newChild
}

func (node *Node) AppendFirstChild(child *Node) {
	child.Detach()
	firstChild := node.firstChild
	node.firstChild = child
	child.nextSibling = firstChild
}

func (node *Node) InsertAfterChild(child, refChild *Node) {
	if refChild == nil {
		// append first 
		node.AppendFirstChild(child)
	} else {
		if refChild.parent != node {
			panic("refchild parent have to be the same as node")
		} else {
			child.Detach()
			nextSibling := refChild.nextSibling
			refChild.nextSibling = child
			child.nextSibling = nextSibling
			child.parent = node
		}
	}
}

// Append a node at the last position of its list of children
func (node *Node) AppendLastChild(child *Node) {
	child.Detach()
	if node.firstChild == nil {
		node.firstChild = child
	} else {
		// cursor is the first child of node
		cursor := node.firstChild
		for cursor.nextSibling != nil {
			cursor = cursor.nextSibling
		}
		// now cursor is the last child of node
		cursor.nextSibling = child
	}
	child.parent = node
}

func (node *Node) Replace(newNode *Node) {
	newNode.Detach()
	newNode.parent = node.parent
	newNode.nextSibling = node.nextSibling
	if node.parent != nil {
		cursor := node.parent.firstChild
		if cursor == node {
			// special case
			node.parent.firstChild = newNode
		} else {
			for cursor.nextSibling != node {
				cursor = cursor.nextSibling
			}
			// cursor is now the previous sibling of node
			cursor.nextSibling = newNode
		}
	}
	node.nextSibling = nil
	node.parent = nil
}

// Is this node has children
func (node *Node) HasChildren() bool {
	return node.firstChild != nil
}

func (node *Node) Clone(deep bool) *Node {
	clone := NewNode(node.element)
	if deep {
		if node.firstChild != nil {
			clone.firstChild = NewNode(node.firstChild.element)
			clone.firstChild.parent = clone
			cloneChild := clone.firstChild
			child := node.firstChild
			for child.nextSibling != nil {
				cloneChild.nextSibling = NewNode(child.nextSibling.element)
				cloneChild.parent = clone
				cloneChild = cloneChild.nextSibling
				child = child.nextSibling
			}
		}
	}
	return clone
}

func (node *Node) RunHierarchy(f func(n *Node)) {
	f(node)
	parent := node.parent
	for parent != nil {
		f(parent)
	}
}

func (node *Node) RunChildren(f func(n *Node)) {

	cursor := node.firstChild
	for cursor != nil {
		f(cursor)
		cursor = node.nextSibling
	}
}

func (node *Node) RunBreadthFirst(f func(n *Node)) {
	var fifo Fifo
	// begin appending to the fifo
	fifo.Add(node)
	var cursor *Node
	for fifo.HasNext() { // No more element in the fifo
		// getting the first element
		cursor = fifo.Next()

		f(cursor)
		cursor = cursor.firstChild
		for cursor != nil {
			fifo.Add(cursor)
			cursor = cursor.nextSibling
		}
	}
}

func (node *Node) RunDeepSuffix(f func(n *Node)) {
	if node.firstChild != nil {
		node.firstChild.runDeepSuffix(f)
	}
	f(node)
}

func (node *Node) runDeepSuffix(f func(n *Node)) {
	if node.firstChild != nil {
		node.firstChild.runDeepSuffix(f)
	}
	f(node)
	if node.nextSibling != nil {
		node.nextSibling.runDeepSuffix(f)
	}
}

func (node *Node) RunDeepPrefix(f func(n *Node)) {
	f(node)
	if node.firstChild != nil {
		node.firstChild.runDeepSuffix(f)
	}
}

func (node *Node) runDeepPrefix(f func(n *Node)) {
	f(node)
	if node.firstChild != nil {
		node.firstChild.runDeepPrefix(f)
	}
	if node.nextSibling != nil {
		node.nextSibling.runDeepPrefix(f)
	}
}
