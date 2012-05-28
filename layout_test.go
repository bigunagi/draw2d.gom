package layout

import (
	"testing"
	"image"
)

func TestTree(t *testing.T) {
	canvasBounds := image.Rect(0, 0, 200, 200)

	node := NewNode("toto")

	node.AppendFirstChild(NewNode("titi"))
	node.AppendFirstChild(NewNode("tata"))
	node.RunChildren(func (childNode *Node) {
		element := childNode.Element()
	})
	
}
