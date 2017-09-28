// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package hashset

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/sets"
)

func assertSetImplementation() {
	var _ sets.Set = (*Set)(nil)
}

// Set holds elements in go's native map
type Set struct {
	items map[interface{}]struct{}
}

var itemExists = struct{}{}

// New instantiates a new empty set
func New() *Set {
	return &Set{items: make(map[interface{}]struct{})}
}

// Add adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Intersect returns whether two sets intersect (i.e. share one or more item).
func (set *Set) Intersect(other sets.Set) bool {
	for k := range set.items {
		if other.Contains(k) {
			return true
		}
	}
	return false
}

// Intersection returns the set of elements that are contained in both initial sets.Set.
func (set *Set) Intersection(other sets.Set) sets.Set {
	inter := New()
	for k := range set.items {
		if other.Contains(k) {
			inter.Add(k)
		}
	}
	return inter
}

// Union returns the set of elements of both sets.
func (set *Set) Union(other sets.Set) sets.Set {
	union := New()
	union.Add(set.items)
	union.Add(other.Values()...)
	return union
}

// Subtract returns the set of elements from the first set that are not in the second set.
func (set *Set) Subtract(other sets.Set) sets.Set {
	diff := New()
	for k := range set.items {
		if !other.Contains(k) {
			diff.Add(k)
		}
	}
	return diff
}

// Empty returns true if set does not contain any elements.
func (set *Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.items = make(map[interface{}]struct{})
}

// Values returns all items in the set.
func (set *Set) Values() []interface{} {
	values := make([]interface{}, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}
