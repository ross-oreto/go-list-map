package maps

import (
	"math/rand"
	"reflect"
	"testing"
)

func perm(n int) []interface{} {
	p := make([]interface{}, n)
	for i, v := range rand.Perm(n) {
		p[i] = v
	}
	return p
}

func randomListMap(n int) *Map {
	_, listmap := randomPermWithListMap(n)
	return listmap
}

func randomPermWithListMap(n int) ([]interface{}, *Map) {
	listmap := New()
	p := perm(n)
	for _, k := range p {
		listmap.Put(k, string(k.(int)+1))
	}
	return p, listmap
}

func TestListMap_Put(t *testing.T) {
	listmap := New()

	listmap.Put(35, "Ross").Put(1, "Test")
	if listmap.Len() != 2 {
		t.Error(listmap.Len(), "list map len should equal 2")
	}

	listmap.PutEntries([]Entry{{Key: 55, Value: "a"}, {Key: 44, Value: "b"}})
	if listmap.Len() != 4 {
		t.Error(listmap.Len(), "list map len should equal 4")
	}

	listmap.PutEntries([]Entry{{Key: 35, Value: "Oreto"}, {Key: 1, Value: "Michael"}})
	if listmap.Len() != 4 {
		t.Error(listmap.Len(), "list map len should equal 4")
	}
}

func TestListMap_Contains(t *testing.T) {
	p, listmap := randomPermWithListMap(1000)

	if !listmap.Contains(900) {
		t.Error("list map should contain 900")
	}
	if !listmap.ContainsAll(p) {
		t.Error("list map should contain all values")
	}
	if !listmap.ContainsAny([]interface{}{900}) {
		t.Error("list map should contain at least one value")
	}
}

func TestListMap_Delete(t *testing.T) {
	p, listmap := randomPermWithListMap(1000)
	listmap.DeleteAll(p)
	if listmap.NotEmpty() {
		t.Error("list map should be empty")
	}
}

func TestListMap_Pop(t *testing.T) {
	listmap := randomListMap(3)
	listmap.Pop()
	listmap.Pop()
	listmap.Pull()
	actual := listmap.Pull()
	if actual != nil {
		t.Error("value of pull should be nil not", actual)
	}
	if listmap.NotEmpty() {
		t.Error("list map should be empty")
	}
}

func TestListMap_Head(t *testing.T) {
	listmap := New()
	listmap.PutEntries([]Entry{{Key: 1, Value: string(1)}, {Key: 2, Value: string(2)}, {Key: 3, Value: string(3)}})

	expect := string(1)
	actual := listmap.Head()
	if actual != expect {
		t.Error(actual, "should equal", expect)
	}

	expect = string(3)
	actual = listmap.Tail()
	if actual != expect {
		t.Error(actual, "should equal", expect)
	}
}

func TestListMap_Entries(t *testing.T) {
	listmap := New()
	entries := []Entry{{Key: 1, Value: string(1)}, {Key: 2, Value: string(2)}, {Key: 3, Value: string(3)}}
	listmap.PutEntries(entries)

	a1 := listmap.Entries()
	if !reflect.DeepEqual(a1, entries) {
		t.Error(a1, "list map entries should equal", entries)
	}

	e2 := []interface{}{1, 2, 3}
	a2 := listmap.Keys()
	if !reflect.DeepEqual(a2, e2) {
		t.Error(a2, "list map entries should equal", e2)
	}

	e2 = []interface{}{string(1), string(2), string(3)}
	a2 = listmap.Values()
	if !reflect.DeepEqual(a2, e2) {
		t.Error(a2, "list map entries should equal", e2)
	}
}

var permutation []interface{}
var listMap *Map

const benchsize = 1000000

func BenchmarkListMap_Put(b *testing.B) {
	permutation = perm(benchsize)
	listMap = New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range permutation {
			listMap.Put(k, string(k.(int)+1))
		}
	}
}

func BenchmarkListMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range permutation {
			listMap.Get(k)
		}
	}
}

func BenchmarkListMap_Entries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, entry := range listMap.Entries() {
			if entry.Key != nil {
			}
		}
	}
}

func BenchmarkListMap_DeleteAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		listMap.DeleteAll(permutation)
	}
}
