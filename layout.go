// Copyright 2012 The draw2d Authors. All rights reserved.
// created: 28/05/2012 by Laurent Le Goff

// gom package
package gom

import ()

type Element interface {
}

/*
type BoxModel int

const (
	Fixed BoxModel = iota
	Flex
	HFlex
	VFlex
	Horizontal
	Vertical
	ConstrainedBox
)

type Element interface {
	Box(node *Node) image.Rectangle
}

func computeBox(constrainedBox image.Rectangle, boxModel BoxModel) image.Rectangle {

}


type FixBox struct {
	box image.Rectangle
}

func (f *FixBox) Box(node *Node) image.Rectangle{
	return f.box
}

type FlexBox struct {
	box image.Rectangle
}

func (f *FlexBox) Box(node *Node) image.Rectangle{
	rect := f.box
	node.RunChildren(func (node *Node) {rect.Union(node.element.Box(node))})
	return rect
}

type HBox struct {
}

func (f *HBox) Box(node *Node) image.Rectangle{
	rect := image.Rect(0, 0, 0, 0)
	node.RunChildren(func (node *Node) {
		box := node.element.Box(node)
		rect.Max.X = rect.Max.X + box.Dx()
	})
	return rect
}

type VBox struct {
}

func (f *VBox) Box(node *Node) image.Rectangle{
	rect := image.Rect(0, 0, 0, 0)
	node.RunChildren(func (node *Node) {
		box := node.element.Box(node)
		rect.Max.Y = rect.Max.Y + box.Dy()
	})
	return rect
}
*/
