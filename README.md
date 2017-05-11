### Go Tree Map

* A map implementation ordered by keys.
* The ordering is done using a backing go tree - ross-oreto/go-tree
* Why not just use a map[] ?
  * Because maps don't guarantee any order while go-tree guarantees natural ordering by the entry key.
  * Note: Keys must be types which are allowed as the key of a standard go map.

#### Basic Usage
```
 import (
     "fmt"
 	"github.com/ross-oreto/go-tree-map"
 )
 
 treeMap := go_tree_map.New()
 treeMap.Put(1, 'Oreto').Put(2, 'Michael').Put(3, 'Ross')
 fmt.Println(treeMap.Values())
```

#### Performance
```
go test -v -benchmem -count 1 -run none -bench .
BenchmarkTreeMap_Put-4                 1        1912420300 ns/op        227584672 B/op   3058928 allocs/op
BenchmarkTreeMap_Get-4                10         117150330 ns/op               0 B/op          0 allocs/op
BenchmarkTreeMap_Entries-4             5         203396580 ns/op        32006144 B/op          1 allocs/op
BenchmarkTreeMap_DeleteAll-4           1        1248692800 ns/op               0 B/op          0 allocs/op
```

- As expected Get operations are really fast, since it's just using the underlying map.
- Put and delete operations are a little slower than a plain btree or map since each insertion/delete must perform the operation on the underlying tree and map.
