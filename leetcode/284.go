//https://leetcode-cn.com/problems/peeking-iterator/

package main

type PeekingIterator struct {
	iter *Iterator
	_hasNext bool
	_next int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, iter.hasNext(), iter.next()}
}

func (this *PeekingIterator) hasNext() bool {
	return this._hasNext;
}

func (this *PeekingIterator) next() int {
	ret := this._next
	this._hasNext = this.iter.hasNext()
	if (this._hasNext){
		this._next = this.iter.next()
	}
	return ret;
}

func (this *PeekingIterator) peek() int {
	return this._next;
}
