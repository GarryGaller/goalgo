//https://golangbyexample.com/go-iterator-design-pattern/
//https://groups.google.com/forum/#!topic/golang-nuts/v6m86sTRbqA
//https://github.com/johnchildren/go-iter


package iterators

import (
    "fmt"
)


type Iterator interface {
    Next() bool            // перемещение на одну позицию вперед в контейнере элементов
    Prev()  bool
    Value() interface{}    // текущий элемент в контейнере
    Index() int
    End()                  // перемещение указателя после последнего элемента контейнера
    Begin()                // перемещение указателя индекса перед начальным элементов контейнера
    Last()  bool           // перемещение указателя на последний элемент контейнера
    First() bool           // перемещение указателя на первый элемент контейнера

}

type Iterable interface {
    Iter() Iterator

}


type AnyIterator struct {
    data    *AnyCollection
    index   int
    size    int
}

func (it *AnyIterator) Value() interface{}  {
    return (*it.data)[it.index]
}

func (it *AnyIterator) Index() int  {
    return it.index
}


func (it *AnyIterator) Next() bool {
    
    if it.index < it.size {
        it.index++
    }
    
    return it.index >= 0 && it.index < it.size
}


func (it *AnyIterator) Prev() bool {
    
    if it.index >= 0 {
        it.index--
    }
    return it.index >= 0 && it.index < it.size
}



func (it *AnyIterator) Begin() {
    it.index = -1
}

func (it *AnyIterator) End() {
    it.index = it.size
}


func (it *AnyIterator) First() bool {
    it.Begin()
    return it.Next()
}


func (it *AnyIterator) Last() bool {
    it.End()
    return it.Prev()
}

type StringIterator struct {
    AnyIterator
    data    *StringCollection
    size int
}

func (it *StringIterator) Value() interface {} {
    return (*it.data)[it.AnyIterator.index]
}

func (it *StringIterator) Index() int  {
    return it.AnyIterator.index
}


func (it *StringIterator) Next() bool {
    
    if it.AnyIterator.index < it.size {
        it.AnyIterator.index++
    }
    
    return (it.AnyIterator.index >= 0 && 
            it.AnyIterator.index < it.AnyIterator.size)
}


func (it *StringIterator) Begin() {
    it.AnyIterator.Begin()
}

type IntIterator struct {
    data    []int
    index   int
}




type AnyCollection []interface{}

type StringCollection []string

func CollectionFromStrings(data []string) AnyCollection {
    collect := make([]interface{},len(data)) 
    for i,v := range data{
        collect[i] = v    
    }
    return collect 
}

func CollectionFromStrings2(data ...interface{}) AnyCollection {
    collect := make([]interface{},len(data)) 
    for i,v := range data{
        collect[i] = v    
    }
    return collect 
}


func CollectionFrom(data ...interface{}) AnyCollection {
    collect := make([]interface{},len(data)) 
    for i,v := range data{
        collect[i] = v    
    }
    return collect 
}


func CollectionFromInts(data []int) AnyCollection {
    collect := make([]interface{},len(data)) 
    for i,v := range data{
        collect[i] = v    
    }
    return collect 
}


func (collect *AnyCollection) Iter() Iterator {
    return &AnyIterator{data: collect, index: -1, size:len(*collect)}
}

func (collect *StringCollection) Iter() Iterator {
    return &StringIterator{data: collect, size:len(*collect )}
}



