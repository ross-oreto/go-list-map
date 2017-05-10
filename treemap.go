package go_tree

import (
	"github.com/ross-oreto/go-tree"
	"fmt"
)

type TreeMap struct {
	entries map[interface{}]interface{}
	keys *tree.Btree
}

type Entry struct {
	Key, Value interface{}
}

func New() *TreeMap { return new(TreeMap).Init() }

func (t *TreeMap) Init() *TreeMap {
	t.entries = make(map[interface{}]interface{})
	if t.keys == nil {
		t.keys = tree.New()
	} else {
		t.keys.Init()
	}
	return t
}

func (t *TreeMap) Put(key interface{}, value interface{}) *TreeMap {
	t.keys.Insert(key)
	t.entries[key] = value
	return t
}

func (t *TreeMap) PutMap(entries map[interface{}]interface{}) *TreeMap {
	for key, value := range entries {
		t.Put(key, value)
	}
	return t
}

func (t *TreeMap) PutEntries(entries []Entry) *TreeMap {
	for _, entry := range entries {
		t.Put(entry.Key, entry.Value)
	}
	return t
}

func (t *TreeMap) Delete(key interface{}) *TreeMap {
	t.keys.Delete(key)
	t.entries[key] = nil
	return t
}

func (t *TreeMap) DeleteAll(keys []interface{}) *TreeMap {
	for _, k := range keys {
		t.Delete(k)
	}
	return t
}

func (t *TreeMap) Keys() []interface{} {
	return t.keys.Values()
}

func (t *TreeMap) String() string {
	return fmt.Sprint(t.Entries())
}

func (t *TreeMap) Values() []interface{} {
	keys := t.keys.Values()
	length := len(keys)
	vals := make([]interface{}, length)
	for i := 0; i < length; i++ {
		vals[i] = t.entries[keys[i]]
	}
	return vals
}

func (t *TreeMap) Entries() []Entry {
	keys := t.keys.Values()
	length := len(keys)
	entries := make([]Entry, length)
	for i := 0; i < length; i++ {
		entries[i] = Entry{Key:keys[i], Value:t.entries[keys[i]]}
	}
	return entries
}

func (t *TreeMap) Map() map[interface{}]interface{} {
	return t.entries
}

func (t *TreeMap) Get(key interface{}) interface{} {
	return t.entries[key]
}

func (t *TreeMap) Contains(key interface{}) bool {
	_, present := t.entries[key]
	return present
}

func (t *TreeMap) ContainsAll(keys []interface{}) bool {
	for _, key := range keys {
		_, present := t.entries[key]
		if !present { return false }
	}
	return true
}

func (t *TreeMap) ContainsAny(keys []interface{}) bool {
	for _, key := range keys {
		_, present := t.entries[key]
		if present { return true }
	}
	return false
}

func (t *TreeMap) Len() int {
	return t.keys.Len()
}

func (t *TreeMap) Head() interface{} {
	return t.entries[t.keys.Head()]
}

func (t *TreeMap) Tail() interface{} {
	return t.entries[t.keys.Tail()]
}

func (t *TreeMap) Pop() interface{} {
	key := t.keys.Pop()
	val := t.entries[key]
	t.entries[key] = nil
	return val
}

func (t *TreeMap) Pull() interface{} {
	key := t.keys.Pull()
	val := t.entries[key]
	t.entries[key] = nil
	return val
}

func (t *TreeMap) Empty() bool {
	return t.keys.Empty()
}

func (t *TreeMap) NotEmpty() bool {
	return t.keys.NotEmpty()
}