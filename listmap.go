package maps

import (
	"container/list"
	"fmt"
)

// Map is an ordered map with some extra useful methods
type Map struct {
	entries map[interface{}]Value
	keys    *list.List
}

// Value represents the value of a map entry
type Value struct {
	val     interface{}
	element *list.Element
}

// Entry represents a map key value pair entry
type Entry struct {
	Key, Value interface{}
}

// New creates a new Map and returns the Map pointer
func New() *Map { return new(Map).Init() }

// Init initializes/clears the Map and returns the Map pointer
func (t *Map) Init() *Map {
	t.entries = make(map[interface{}]Value)
	if t.keys == nil {
		t.keys = list.New()
	} else {
		t.keys.Init()
	}
	return t
}

// Put enters a new key value pair into the map and returns the Map pointer
func (t *Map) Put(key interface{}, value interface{}) *Map {
	if v, ok := t.entries[key]; ok {
		t.entries[key] = Value{value, v.element}
	} else {
		t.entries[key] = Value{value, t.keys.PushBack(key)}
	}
	return t
}

// PutEntries enters all entries into the Map and returns the Map pointer
func (t *Map) PutEntries(entries []Entry) *Map {
	for _, entry := range entries {
		t.Put(entry.Key, entry.Value)
	}
	return t
}

// Delete deletes the key from the map and returns the Map pointer
func (t *Map) Delete(key interface{}) *Map {
	if _, ok := t.entries[key]; ok {
		t.keys.Remove(t.entries[key].element)
		delete(t.entries, key)
	}
	return t
}

// DeleteAll deletes all the keys from the map and returns the Map pointer
func (t *Map) DeleteAll(keys []interface{}) *Map {
	for _, k := range keys {
		t.Delete(k)
	}
	return t
}

// Keys creates and returns a slice of all the map keys
func (t *Map) Keys() []interface{} {
	keys := make([]interface{}, t.keys.Len())
	for e, i := t.keys.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		keys[i] = e.Value
	}
	return keys
}

// String returns a string representing the map entries
func (t *Map) String() string {
	return fmt.Sprint(t.Entries())
}

// Values creates and returns a slice of all the map values
func (t *Map) Values() []interface{} {
	vals := make([]interface{}, t.keys.Len())
	for e, i := t.keys.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		vals[i] = t.entries[e.Value].val
	}
	return vals
}

// Entries creates and returns a slice of all the map key value pair entries
func (t *Map) Entries() []Entry {
	entries := make([]Entry, t.keys.Len())
	for e, i := t.keys.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		entries[i] = Entry{Key: e.Value, Value: t.entries[e.Value].val}
	}
	return entries
}

// Get gets and returns the value for the specified search key
func (t *Map) Get(key interface{}) interface{} {
	if val, ok := t.entries[key]; ok {
		return val.val
	}
	return nil
}

// Contains returns true if the map contains the key
func (t *Map) Contains(key interface{}) bool {
	_, ok := t.entries[key]
	return ok
}

// ContainsAll returns true if the map contains all the keys
func (t *Map) ContainsAll(keys []interface{}) bool {
	for _, key := range keys {
		if !t.Contains(key) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the map contains any of the keys
func (t *Map) ContainsAny(keys []interface{}) bool {
	for _, key := range keys {
		if t.Contains(key) {
			return true
		}
	}
	return false
}

// Len returns the number of map entries
func (t *Map) Len() int {
	return t.keys.Len()
}

// Head returns the first value of the ordered map
func (t *Map) Head() interface{} {
	if t.NotEmpty() {
		return t.entries[t.keys.Front().Value].val
	}
	return nil
}

// Tail returns the last value of the ordered map
func (t *Map) Tail() interface{} {
	if t.NotEmpty() {
		return t.entries[t.keys.Back().Value].val
	}
	return nil
}

// Pop deletes the last map entry and returns its value
func (t *Map) Pop() interface{} {
	if t.NotEmpty() {
		key := t.keys.Remove(t.keys.Back())
		val, ok := t.entries[key]
		delete(t.entries, key)
		if ok {
			return val.val
		}
	}
	return nil
}

// Pull deletes the first map entry and returns its value
func (t *Map) Pull() interface{} {
	if t.NotEmpty() {
		key := t.keys.Remove(t.keys.Front())
		val, ok := t.entries[key]
		delete(t.entries, key)
		if ok {
			return val.val
		}
	}
	return nil
}

// Empty returns true if the Map is empty
func (t *Map) Empty() bool {
	return t.keys.Len() == 0
}

// NotEmpty returns true if the Map is not empty
func (t *Map) NotEmpty() bool {
	return t.keys.Len() > 0
}
