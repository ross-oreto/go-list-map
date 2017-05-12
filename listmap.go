package go_list_map

import (
	"fmt"
	"container/list"
)

type ListMap struct {
	entries map[interface{}]Value
	keys *list.List
}

type Value struct {
	val interface{}
	element *list.Element
}

type Entry struct {
	Key, Value interface{}
}

func New() *ListMap { return new(ListMap).Init() }

func (t *ListMap) Init() *ListMap {
	t.entries = make(map[interface{}]Value)
	if t.keys == nil {
		t.keys = list.New()
	} else {
		t.keys.Init()
	}
	return t
}

func (t *ListMap) Put(key interface{}, value interface{}) *ListMap {
	if v, ok := t.entries[key]; ok {
		t.entries[key] = Value{value, v.element}
	} else {
		t.entries[key] = Value{value, t.keys.PushBack(key)}
	}
	return t
}

func (t *ListMap) PutMap(entries map[interface{}]interface{}) *ListMap {
	for key, value := range entries {
		t.Put(key, value)
	}
	return t
}

func (t *ListMap) PutEntries(entries []Entry) *ListMap {
	for _, entry := range entries {
		t.Put(entry.Key, entry.Value)
	}
	return t
}

func (t *ListMap) Delete(key interface{}) *ListMap {
	if _, ok := t.entries[key]; ok {
		t.keys.Remove(t.entries[key].element)
		delete(t.entries, key)
	}
	return t
}

func (t *ListMap) DeleteAll(keys []interface{}) *ListMap {
	for _, k := range keys {
		t.Delete(k)
	}
	return t
}

func (t *ListMap) Keys() []interface{} {
	keys := make([]interface{}, t.keys.Len())
	for e, i := t.keys.Front(), 0; e != nil; e, i = e.Next(), i + 1 {
		keys[i] = e.Value
	}
	return keys
}

func (t *ListMap) String() string {
	return fmt.Sprint(t.Entries())
}

func (t *ListMap) Values() []interface{} {
	vals := make([]interface{}, t.keys.Len())
	for e, i := t.keys.Front(), 0; e != nil; e, i = e.Next(), i + 1 {
		vals[i] = t.entries[e.Value].val
	}
	return vals
}

func (t *ListMap) Entries() []Entry {
	entries := make([]Entry, t.keys.Len())
	for e, i := t.keys.Front(), 0; e != nil; e, i = e.Next(), i + 1 {
		entries[i] = Entry{Key:e.Value, Value:t.entries[e.Value].val}
	}
	return entries
}

func (t *ListMap) Get(key interface{}) interface{} {
	if val, ok := t.entries[key]; ok {
		return val.val
	}
	return nil
}

func (t *ListMap) Contains(key interface{}) bool {
	_, ok := t.entries[key]
	return ok
}

func (t *ListMap) ContainsAll(keys []interface{}) bool {
	for _, key := range keys {
		if !t.Contains(key) { return false }
	}
	return true
}

func (t *ListMap) ContainsAny(keys []interface{}) bool {
	for _, key := range keys {
		if t.Contains(key) { return true }
	}
	return false
}

func (t *ListMap) Len() int {
	return t.keys.Len()
}

func (t *ListMap) Head() interface{} {
	if t.NotEmpty() { return t.entries[t.keys.Front().Value].val }
	return nil
}

func (t *ListMap) Tail() interface{} {
	if t.NotEmpty() { return t.entries[t.keys.Back().Value].val }
	return nil
}

func (t *ListMap) Pop() interface{} {
	if t.NotEmpty() {
		key := t.keys.Remove(t.keys.Back())
		val, ok := t.entries[key]
		delete(t.entries, key)
		if ok { return val.val }
	}
	return nil
}

func (t *ListMap) Pull() interface{} {
	if t.NotEmpty() {
		key := t.keys.Remove(t.keys.Front())
		val, ok := t.entries[key]
		delete(t.entries, key)
		if ok { return val.val }
	}
	return nil
}

func (t *ListMap) Empty() bool {
	return t.keys.Len() == 0
}

func (t *ListMap) NotEmpty() bool {
	return t.keys.Len() > 0
}