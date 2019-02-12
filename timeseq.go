/*
Package timeseq provide a sorted time sequence.
*/
package timeseq

import (
	"sort"
	"time"
)

// Elem is an element in the sequence
type Elem struct {
	Time  time.Time
	Value interface{}
}

// TimeSeq is a time sequence
type TimeSeq []*Elem

// New creates a new TimeSeq. initialCap is the initial capacity
func New(initialCap int) *TimeSeq {
	ts := TimeSeq(make([]*Elem, 0, initialCap))
	return &ts
}

// Add adds an value at time point
func (ts *TimeSeq) Add(t time.Time, value interface{}) {
	fn := func(i int) bool {
		return (*ts)[i].Time.After(t)
	}
	i := sort.Search(len(*ts), fn)
	*ts = append(*ts, nil)
	// shift slice right of i
	copy((*ts)[i+1:], (*ts)[i:])
	(*ts)[i] = &Elem{t, value}
}

// Len returns the number of elements in ts
func (ts *TimeSeq) Len() int {
	return len(*ts)
}

// Slice return all elements in (closed) time interval
func (ts *TimeSeq) Slice(start, end time.Time) []*Elem {
	fn := func(i int) bool {
		return (*ts)[i].Time.After(start)
	}
	s := sort.Search(len(*ts), fn)
	// Include equal times
	for s > 0 && start.After((*ts)[s].Time) {
		s--
	}

	fn = func(i int) bool {
		return (*ts)[i].Time.After(end)
	}
	e := sort.Search(len(*ts), fn)
	return (*ts)[s:e]
}

// Drop will drop all records until time
func (ts *TimeSeq) Drop(until time.Time) {
	fn := func(i int) bool {
		return (*ts)[i].Time.After(until)
	}
	i := sort.Search(len(*ts), fn)
	*ts = (*ts)[i:]
}
