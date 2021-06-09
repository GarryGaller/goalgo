package iterators

import (
    "fmt"
    "testing"
)

func Test(t *testing.T) {

    strSlice := []string{"zero", "one", "two","three", "four", "five"}
    
    data := &AnyCollection{1,2,3,4,5} 
    
    it := data.Iter()
    for it.Next(){ 
        fmt.Println(it.Value().(int))
    }
    
    collection := &AnyCollection{"zero", "one", "two","three", "four", "five"} 
    it = collection.Iter()
    for it.Next(){ 
        fmt.Println(it.Value().(string))
    }
    
    
    
    strColl := StringCollection(strSlice) // приведение к типу
    strIt := strColl.Iter()
    for strIt.Next() { 
       fmt.Println(it.Index(), it.Value().(string))
    }
    
    
    
    //data2 := CollectionFromStrings(strSlice)
    data2 := CollectionFrom("zero", "one", "two","three", "four")
    fmt.Printf("%#v\n", data2)
    it = data2.Iter()
    for it.Next(){ 
        fmt.Println(it.Index(), it.Value().(string))
    }
    
    it.Begin()
    for it.Next(){ 
        fmt.Println(it.Index(), it.Value().(string))
    }
    
    it.Last()
    fmt.Println(it.Index(), it.Value().(string))
    it.First()
    fmt.Println(it.Index(), it.Value().(string))
    
}
