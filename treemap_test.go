package go_tree_map

import (
	"testing"
	"math/rand"
	"reflect"
)

func perm(n int) []interface{} {
	p := make([]interface{}, n)
	for i, v := range rand.Perm(n) {
		p[i] = v
	}
	return p
}

func randomTreeMap(n int) *TreeMap {
	_, treemap := randomPermWithTreeMap(n)
	return treemap
}

func randomPermWithTreeMap(n int) ([]interface{}, *TreeMap) {
	treemap := New()
	p := perm(n)
	for _, k := range p { treemap.Put(k, string(k.(int) + 1)) }
	return p, treemap
}

func TestTreeMap_Put(t *testing.T) {
	treemap := New()

	treemap.Put(35, "Ross").Put(1, "Test")
	if treemap.Len() != 2 { t.Error(treemap.Len(),"tree map len should equal 2") }

	treemap.PutMap(map[interface{}]interface{}{55: "a", 44: "b" })
	if treemap.Len() != 4 { t.Error(treemap.Len(),"tree map len should equal 4") }

	treemap.PutEntries([]Entry{{Key:35, Value:"Oreto"},{Key:1, Value:"Michael"}})
	if treemap.Len() != 4 { t.Error(treemap.Len(),"tree map len should equal 4") }
}

func TestTreeMap_Contains(t *testing.T) {
	p, treemap := randomPermWithTreeMap(1000)

	if !treemap.Contains(900) {  t.Error("tree map should contain 900") }
	if !treemap.ContainsAll(p) { t.Error("tree map should contain all values") }
	if !treemap.ContainsAny([]interface{}{ 900 }) { t.Error("tree map should contain at least one value")  }
}

func TestTreeMap_Delete(t *testing.T) {
	p, treemap := randomPermWithTreeMap(1000)
	treemap.DeleteAll(p)
	if treemap.NotEmpty() {  t.Error("tree map should be empty")}
}

func TestTreeMap_Pop(t *testing.T) {
	treemap := randomTreeMap(3)
	treemap.Pop()
	treemap.Pop()
	treemap.Pull()
	actual := treemap.Pull()
	if actual != nil { t.Error("value of pull should be nil not", actual) }
	if treemap.NotEmpty() {  t.Error("tree map should be empty") }
}

func TestTreeMap_Head(t *testing.T) {
	treemap := randomTreeMap(3)

	expect := string(1)
	actual := treemap.Head()
	if actual != expect { t.Error(actual, "should equal", expect)}

	expect = string(3)
	actual = treemap.Tail()
	if actual != expect { t.Error(actual, "should equal", expect)}
}

func TestTreeMap_Entries(t *testing.T) {
	treemap := randomTreeMap(3)

	e1 := []Entry{{Key:0, Value:string(1)},{Key:1, Value:string(2)},{Key:2, Value:string(3)}}
	a1 := treemap.Entries()
	if !reflect.DeepEqual(a1, e1) { t.Error(a1, "tree map entries should equal", e1)}

	e2 := []interface{}{0, 1, 2}
	a2 := treemap.Keys()
	if !reflect.DeepEqual(a2, e2) { t.Error(a2, "tree map entries should equal", e2)}

	e2 = []interface{}{string(1), string(2), string(3)}
	a2 = treemap.Values()
	if !reflect.DeepEqual(a2, e2) { t.Error(a2, "tree map entries should equal", e2)}
}

var permutation []interface{}
var treeMap *TreeMap

const benchsize = 1000000

func BenchmarkTreeMap_Put(b *testing.B) {
	permutation = perm(benchsize)
	treeMap = New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range permutation { treeMap.Put(k, string(k.(int) + 1)) }
	}
}

func BenchmarkTreeMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range permutation {
			treeMap.Get(k)
		}
	}
}

func BenchmarkTreeMap_Entries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, entry := range treeMap.Entries() {
			if entry.Key != nil {}
		}
	}
}

func BenchmarkTreeMap_DeleteAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		treeMap.DeleteAll(permutation)
	}
}