### Go List Map

* A ordered map implementation ordered by insertion order.
* The ordering is done using a backing list - "container/list"
* Why not just use a map[] ?
  * Because maps don't guarantee any order while go-list-map guarantees ordering
  * Note: Keys must be types which are allowed as the key of a standard go map.

#### Basic Usage
```
 import (
     "fmt"
 	"github.com/ross-oreto/go-list-map"
 )
 
 listMap := go_list_map.New()
 listMap.Put(1, 'Oreto').Put(2, 'Michael').Put(3, 'Ross')
 fmt.Println(listMap.Values())
```

#### Performance
```
go test -v -benchmem -count 1 -run none -bench .
BenchmarkListMap_Put-4                 2         529298900 ns/op        142446192 B/op   2529443 allocs/op
BenchmarkListMap_Get-4                10         114120630 ns/op               0 B/op          0 allocs/op
BenchmarkListMap_Entries-4            10         147854010 ns/op        32006144 B/op          1 allocs/op
BenchmarkListMap_DeleteAll-4         300           3884690 ns/op               0 B/op          0 allocs/op
```
