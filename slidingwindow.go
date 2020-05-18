package slidingwindow

import (
	"errors"
	"fmt"
)

// Element is a wrapper
// TODO make it a linked list maybe?
type Element interface{}

// SlidingWindow represents a sliding window structure.
type SlidingWindow struct {
	size  int
	index int
	len   int
	queue []Element
}

// New returns an empty sliding window.
func New(size int) (*SlidingWindow, error) {
	if size <= 0 {
		return nil, errors.New("Size must be > 0")
	}
	mw := SlidingWindow{}
	mw.size = size
	mw.queue = make([]Element, mw.size)
	mw.index = 0
	return &mw, nil
}

// Add adds an element to the sliding window.
func (mw *SlidingWindow) Add(elem Element) {
	mw.queue[mw.index] = elem
	mw.index++
	if mw.index == mw.size {
		mw.index = 0
	}
	if mw.len < mw.size {
		mw.len++
	}
}

func (mw *SlidingWindow) start() int {
	return (mw.index - mw.len + mw.size) % mw.size
}

// Contains returns true if elem is contained in the sliding window, false
// otherwise.
func (mw *SlidingWindow) Contains(elem Element) bool {
	for i := mw.start(); i < mw.start()+mw.len; i++ {
		if mw.queue[i%mw.size] == elem {
			return true
		}
	}
	return false
}

// String returns a string representing the sliding window
func (mw *SlidingWindow) String() string {
	if mw.len == 0 {
		return "[]"
	}

	s := "["
	for i := mw.start(); i < mw.start()+mw.len; i++ {
		s += fmt.Sprint(mw.queue[i%mw.size], " ")
	}
	s = s[:len(s)-1]
	s += fmt.Sprint("]")
	return s
}

// Len returns the length of the sliding window.
func (mw *SlidingWindow) Len() int {
	return mw.len
}
