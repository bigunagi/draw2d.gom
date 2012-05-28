// Copyright 2012 The draw2d Authors. All rights reserved.
// created: 28/05/2012 by Laurent Le Goff

package gom

import (
	"image"
	"testing"
)

func TestTree(t *testing.T) {
	canvasBounds := image.Rect(0, 0, 200, 200)

	node := NewNode("toto")

	node.AppendFirstChild(NewNode("titi"))
	node.AppendFirstChild(NewNode("tata"))
	node.RunChildren(func(childNode *Node) {
		element := childNode.Element()
	})

}
