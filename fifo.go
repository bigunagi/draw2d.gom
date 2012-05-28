// Copyright 2012 The draw2d Authors. All rights reserved.
// created: 28/05/2012 by Laurent Le Goff

package gom

type Fifo struct {
	nodes []*Node
	ri    int
	wi    int
}

func (f *Fifo) Add(n *Node) {
	if f.wi == f.ri {
		// Buffer is full, allocate something bigger
		tmp := make([]*Node, len(f.nodes)*2+2)
		// realign elements in the slice
		copy(tmp, f.nodes[f.ri:len(f.nodes)])
		copy(tmp[f.ri:], f.nodes[0:f.ri])
		f.nodes = tmp
	}
	f.nodes[f.wi] = n
	f.wi = (f.wi + 1) % len(f.nodes) // increment writing index
}

func (f *Fifo) HasNext() bool {
	return f.ri != f.wi
}

func (f *Fifo) Next() *Node {
	if f.ri == f.wi {
		// buffer is empty
		return nil
	}
	next := f.nodes[f.ri]
	f.nodes[f.ri] = nil
	f.ri = (f.ri + 1) % len(f.nodes) // increment reading index
	return next
}
