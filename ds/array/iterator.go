package array

import (
	. "github.com/liyue201/gostl/utils/iterator"
)

//ArrayIterator is a RandomAccessIterator
var _ RandomAccessIterator = (*ArrayIterator)(nil)

type ArrayIterator struct {
	array    *Array
	position int
}

func (this *ArrayIterator) IsValid() bool {
	if this.position >= 0 && this.position < this.array.Size() {
		return true
	}
	return false
}

func (this *ArrayIterator) Value() interface{} {
	return this.array.At(this.position)
}

func (this *ArrayIterator) SetValue(val interface{}) error {
	return this.array.Set(this.position, val)
}

func (this *ArrayIterator) Next() ConstIterator {
	if this.position < this.array.Size() {
		this.position++
	}
	return this
}

func (this *ArrayIterator) Prev() ConstBidIterator {
	if this.position >= 0 {
		this.position--
	}
	return this
}

func (this *ArrayIterator) Clone() ConstIterator {
	return &ArrayIterator{array: this.array, position: this.position}
}

func (this *ArrayIterator) IteratorAt(position int) RandomAccessIterator {
	return &ArrayIterator{array: this.array, position: position}
}

func (this *ArrayIterator) Position() int {
	return this.position
}

func (this *ArrayIterator) Equal(other ConstIterator) bool {
	otherIter, ok := other.(*ArrayIterator)
	if !ok {
		return false
	}
	if otherIter.array == this.array && otherIter.position == this.position {
		return true
	}
	return false
}
