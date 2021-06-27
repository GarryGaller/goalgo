package main

import (
    "fmt"
    "strings"
    "sort"
    "math"
    //"math/rand"
    
    lane "gopkg.in/oleiade/lane.v1" //
    //deq "github.com/gammazero/deque"
    
    //"github.com/golang/groupcache"                   //lru\consistenthash                                                                                                                                                                      10K stars
    //"github.com/Workiva/go-datastructures"          // fibheap\queue\set\skiplist\btree\bitarray\augmentedtree\rtree\rangetree\tree-avl\hashmap\trie\graph                                                                                     6K stars
   
   // по числу звездочек на гитхабе самая мощная библиотека - 10k звезд;
   //"github.com/emirpasic/gods"                       //SinglyLinkedList\DoublyLinkedList\HashSet\TreeSet\\LinkedHashSet\LinkedListStack\ArrayStack\HashMap\TreeMap\LinkedHashMap\HashBidiMap\TreeBidiMap\RedBlackTree\AVLTree\BTree\BinaryHeap  10K stars

    
    "github.com/emirpasic/gods/maps/hashbidimap"
    
 
    //"github.com/emirpasic/gods/sets/hashset"        //A set backed by a hash table (actually a Go's map): Add/Remove/Contains/Empty/Size/Values/Clear
    //"github.com/emirpasic/gods/sets/treeset"        //A set backed by a red-black tree to keep the elements ordered with respect to the comparator.
   
   // "github.com/Workiva/go-datastructures/set"      //simple unordered set implemented with a map. Add/Remove/Exists/All
    "github.com/xtgo/set"                             // unordered set implemented with a map: Diff/SymDiff/Inter/Union/IsSub/IsSuper
    //"github.com/goware/set"               // основан на xtgo/set 
    
    // 2k звезд на гитхабе
    //"github.com/deckarep/golang-set"     // unordered set implemented with a map:                        Add/Remove/Pop/Contains/Difference/SymmetricDifference/Intersect/Union/IsSubSet/IsSuperSet/Cardinality/CartesianProduct/Each/Iter  etc.
    
    // same as golang-set - можно использовать не интерфейсные типы
    "github.com/zoumo/goset"               // unordered set implemented with a map (почти как golang-set): Add/Remove/Pop/Contains/Diff/SymmetricDiff/Intersect/Unite/IsSubSetOf/IsSuperSetOf  etc.
    
    "github.com/scylladb/go-set/strset"    // unordered set implemented with a map:                        Add/Remove/Pop/Has/Difference/SymmetricDifference/Intersection/Union/IsSubSet/IsSuperSet
     
    
     //"github.com/liyue201/gostl/algorithm"   //     sort(quick_sort)\stable_sort(merge_sort)\binary_search\lower_bound\upper_bound\next_permutation\nth_element\swap\reverse\count\count_if\find\find_if
    //"github.com/liyue201/gostl/ds"           //     slice\array\vector\list\deque\queue\priority_queue\stack\rbtree(red_black_tree)\map\multimap\set\multiset\bitmap\bloom_filter\hamt(hash_array_mapped_trie)\ketama\skiplist 

    
     // в типе set не реализованы базовые diff\union\inter  операции
     // "github.com/emirpasic/gods/hashset"   Agg/Remove/Contains/Empty/Clear/Size
     
     // "github.com/itrabbit/go-stp"
    
    )
    
func deque_example() {    
    // Let's create a new deque data structure
    var deque *lane.Deque = lane.NewDeque()

    // And push some content into it using the Append
    // and Prepend methods
    deque.Append("easy as")
    deque.Prepend("123")
    deque.Append("do re mi")
    deque.Prepend("abc")

    // Now let's take a look at what are the first and
    // last element stored in the Deque
    firstValue := deque.First()
    lastValue := deque.Last()
    fmt.Println(firstValue) // "abc"
    fmt.Println(lastValue)  // 1

    // Okay now let's play with the Pop and Shift
    // methods to bring the song words together
    var jacksonFive []string = make([]string, deque.Size())

    for i := 0; i < len(jacksonFive); i++ {
        value := deque.Shift()
        jacksonFive[i] = value.(string)
    }

    // abc 123 easy as do re mi
    fmt.Println(strings.Join(jacksonFive, " "))
    // Let's create a new musical quartet
    quartet := lane.NewCappedDeque(4)

    // List of young hopeful musicians
    musicians := []string{"John", "Paul", "George", "Ringo", "Stuart"}

    // Add as many of them to the band as we can.
    for _, name := range musicians {
        if quartet.Append(name) {
            fmt.Printf("%s is in the band!\n", name)
        } else {
            fmt.Printf("Sorry - %s is not in the band.\n", name)
        }
    }

    // Assemble our new rock sensation
    var beatles = make([]string, quartet.Size())

    for i := 0; i < len(beatles); i++ {
        beatles[i] = quartet.Shift().(string)
    }

    fmt.Println("The Beatles are:", strings.Join(beatles, ", "))    
}


    
    
func test(){
  // словарь в котором не только ключи должны быть уникальными, но и значения
    m := hashbidimap.New() // empty
    m.Put(1, "x")          // 1->x
    m.Put(3, "b")          // 1->x, 3->b (random order)
    m.Put(1, "a")          // 1->a, 3->b (random order)
    m.Put(2, "b")          // 1->a, 2->b (random order)
    _, _ = m.GetKey("a")   // 1, true
    _, _ = m.Get(2)        // b, true
    _, _ = m.Get(3)        // nil, false
    fmt.Printf("%#v\n",m.Values())         // []interface {}{"a", "b"} (random order)
    fmt.Printf("%#v\n",m.Keys())           // []interface {}{1, 2} (random order)
    m.Remove(1)            // 2->b
    m.Clear()              // empty
    m.Empty()              // true
    m.Size()               // 0
}


type Args struct{
    x,y int
    str string
    str2 string
    flag bool
}


func test2(args *Args){
    fmt.Printf("%#v\n",args)

}

func test3(args ...interface{}) {
     fmt.Printf("%#v\n",args)
    fmt.Printf("%v %v\n",args[0].(string),args[1].(int))

}

func deduplicate() {
    in := []int{3, 2, 1, 4, 3, 2, 1, 4, 1}
    fmt.Println(in)
    sort.Ints(in)
    fmt.Println(in)
    j := 0
    for i := 1; i < len(in); i++ {
        if in[j] == in[i] {
            continue
        }
        j++
        in[j] = in[i]
        fmt.Println(in)
    }
    in = in[:j+1]
    fmt.Println(in)
}

func countMaxElements() {
    
    vector :=  []int{1,2,3,4,5,6,7,8,9,10,10,10}
    sort.Sort(sort.Reverse(sort.IntSlice(vector)))
    
    count := 1
    max := vector[0]
    
    for i:= 1; i < len(vector);i++ {
        if vector[i] == max {
            count+=1
        } else {
            break
        }       
    }
    fmt.Printf("%#v\n", vector)
    fmt.Printf("%d %d\n",max, count) 
}


func countMaxElements2() {
    
    vector :=  []int{1,2,3,4,5,6,7,8,9,10,10,10}
    count := 0
    max := math.Inf(-1)
    
    for i:= 0; i < len(vector);i++ {
        
        switch {
        case float64(vector[i]) == max:
            count++
        case float64(vector[i]) > max:
            max = float64(vector[i])
            count = 1
        }
    }
    fmt.Printf("%#v\n", vector)
    fmt.Printf("%f %d\n",max, count) 
}


func main() {
    //test()
    
    old:= []string{"1","2"}
    test:= append(old,"10")
    fmt.Println(test[1:len(test)])
    
    dict_query := map[string][]string{
        "languageisocode":append(old,"10"),
    }
    
    fmt.Printf("%v\n",dict_query)
    
    args := Args{x:1,y:2, str:"hello"}
    test2(&args)
    
    test3("hello",1)
    
    
    
    data0:= sort.IntSlice{5,5,5,5,5,2,6}
    data1:= sort.IntSlice{1,3,4,5}
    sort.Sort(data0)
    sort.Sort(data1)
    
    n0 := set.Uniq(data0)
    n1 := set.Uniq(data1)
    pivot := len(data0[:n0])
    data_all := append(data0[:n0], data1[:n1]...)
    
    fmt.Printf("data_all: %v\n",data_all)
    fmt.Printf("data0 uniq: %v\n",data0[:n0])
    fmt.Printf("data1 uniq: %v\n",data1[:n1])
    
    n3 := set.Diff(data_all, pivot)
    diff0 := data_all[:n3]
    fmt.Printf("diff: %v\n",diff0)   // [2]
    fmt.Printf("diff: %v\n",data_all) 
    
    n4 := set.SymDiff(data_all, pivot) // [1 2 3 4]
    diff1 := data_all[:n4]
    fmt.Printf("symdiff: %v\n",diff1) 
    
    fmt.Println("---------------")
    
    data := sort.IntSlice{3, 5, 7, 5}
    fmt.Printf("data: %v\n",data)
    
    sort.Sort(data)
    n := set.Uniq(data)
    uniq := data[:n]
    fmt.Printf("uniq: %v\n",uniq)
    fmt.Printf("uniq: %v\n",set.Ints([]int{3, 5, 7, 5}))
    
    pivot = len(data)
   
    data2 :=sort.IntSlice{7,5,8}
    fmt.Printf("data2:%v\n",data2)
    
    sort.Sort(data2)
    data_ := append(data, data2...)
   
    
    size := set.Inter(data_, pivot)
    inter := data_[:size]
    fmt.Println("inter:", inter) 
    
    size = set.Diff(data_, pivot)
    diff := data_[:size]
    fmt.Println("diff:", diff) 
     
    
    size = set.SymDiff(data_, pivot)
    symdiff := data_[:size]
    fmt.Println("symdiff:", symdiff) 
    
    type dict map[string]interface{}
   
    d:=dict{"x":1,"y":2}
   
    fmt.Printf("%#v\n",&d)
    d["z"]=3
   
    fmt.Printf("%#v\n",&d)
   
    //var m map[string]int
    //m["one"] = 1 // ошибка
    
    m := make(map[string]int)
    m["one"] = 1 
    
    fmt.Printf("%#v\n",&m)
    
    mm:= map[string]int{"one":1}
    mm["two"] = 2
    mm["three"] = 3
    fmt.Printf("%#v\n",&mm)
    
    
    //fmt.Printf("%#v\n", rand.Perm(100)) 
    
    set0 := goset.NewSet(1, 2, 3, 4,5,6,7,8,9,10)
    fmt.Printf("%v\n",set0)
    // or
    set1 := goset.NewSetFrom([]int{1,2,3,4,5})
    fmt.Printf("%v\n",set1)
    // or
    set2 := goset.NewSet(1,2,3,4,5,6,7)
    fmt.Printf("%v\n",set2)
    // or
    set3 := goset.NewSetFrom([]int{2,3,4,5,6,7,11})
    fmt.Printf("%v\n",set3)
    
    fmt.Println("-------------------")
    
    fmt.Printf("%v\n",set0.Diff(set1)) //Set[6 7 8 10 9]
    fmt.Printf("%v\n",set0.Diff(set2)) //Set[8 10 9]
    fmt.Printf("%v\n",set0.Diff(set3)) //Set[9 1 8 10]
    
    fmt.Printf("%v\n",set0.SymmetricDiff(set1)) //Set[8 9 10 6 7]
    fmt.Printf("%v\n",set0.SymmetricDiff(set2)) //Set[8 10 9]
    fmt.Printf("%v\n",set0.SymmetricDiff(set3)) //Set[1 10 8 9 11]
      
        
    fmt.Printf("%v\n",set0.Intersect(set1)) //Set[2 3 4 5 1]
    fmt.Printf("%v\n",set0.Intersect(set2)) //Set[5 6 7 1 2 3 4]
    fmt.Printf("%v\n",set0.Intersect(set3)) //Set[6 7 2 3 4 5]
    
     
    s1 := strset.New("entry1", "entry2")
    s2 := strset.New("entry2", "entry3")
    s3 := strset.Intersection(s1, s2)
    fmt.Printf("%v\n",s3)
    
    countMaxElements()
    countMaxElements2()
    fmt.Println("----deduplicate----------")
    deduplicate()
    
    
}
