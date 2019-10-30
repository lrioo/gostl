package rbtree

import . "github.com/liyue201/gostl/utils/iterator"

type RbTreeIterator struct {
	node *Node
}

func NewIterator(node *Node) *RbTreeIterator {
	return &RbTreeIterator{node: node}
}

func (this *RbTreeIterator) IsValid() bool {
	if this.node != nil {
		return true
	}
	return false
}

func (this *RbTreeIterator) Next() ConstIterator {
	if this.IsValid() {
		this.node = this.node.Next()
	}
	return this
}

func (this *RbTreeIterator) Prev() ConstBidIterator {
	if this.IsValid() {
		this.node = this.node.Prev()
	}
	return this
}

func (this *RbTreeIterator) Key() interface{} {
	return this.node.Key()
}

func (this *RbTreeIterator) Value() interface{} {
	return this.node.Value()
}

func (this *RbTreeIterator) SetValue(val interface{}) error {
	this.node.SetValue(val)
	return nil
}

func (this *RbTreeIterator) Clone() ConstIterator {
	return NewIterator(this.node)
}

func (this *RbTreeIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*RbTreeIterator)
	if !ok {
		return false
	}
	if otherIter.node == this.node {
		return true
	}
	return false
}

