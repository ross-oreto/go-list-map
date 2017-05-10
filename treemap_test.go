package go_tree

import (
	"testing"
)

func TestTreeMap_Put(t *testing.T) {
	treemap := New()

	treemap.Put(35, "Ross").Put(1, "Test")
	if treemap.Len() != 2 { t.Error("tree map len should equal 2") }

	treemap.PutMap(map[interface{}]interface{}{55: "a", 44: "b" })
	if treemap.Len() != 4 { t.Error("tree map len should equal 4") }

	treemap.PutEntries([]Entry{{Key:35, Value:"Oreto"},{Key:1, Value:"Michael"}})
	if treemap.Len() != 4 { t.Error("tree map len should equal 4") }
}