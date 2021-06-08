package set


type void struct{}
type Set map[interface{}]void

func New(values ...interface{}) *Set {
    set := make(Set, len(values))

    for _, v := range values {
        set.Add(v)
    }

    return &set
}

func (set *Set) Add(v interface{}) {
    (*set)[v] = void{}
}

func (set *Set) Contains(v interface{}) (ok bool) {
    _, ok = (*set)[v]
    return
}

func (set *Set) Delete(v interface{}) {
    delete(*set, v)
}

func (set *Set) Size() int {
    return len(*set)
}

func (set *Set) IsEmpty() bool {
    return set.Card() == 0
}

func (set *Set) Card() int {
    return len(*set)
}

func (set *Set) Inter(other *Set) (inter *Set) {

    inter = New()
    if set.IsEmpty() || other.IsEmpty() {
        return
    }
    // loop over smaller set
    if set.Card() < other.Card() {
        for v := range *set {
            if other.Contains(v) {
                inter.Add(v)
            }
        }
    } else {
        for v := range *other {
            if set.Contains(v) {
                inter.Add(v)
            }
        }
    }

    return
}

func (set *Set) Diff(other *Set) (diff *Set) {
    diff = New()
    for v := range *set {
        if !other.Contains(v) {
            diff.Add(v)
        }
    }
    return
}
 

func (set *Set) SymDiff(other *Set) (diff *Set) {

    diff = set.Diff(other)
    diff2 := other.Diff(set)
    diff.Union(diff2)
    return
}

func (set *Set) Union(other *Set) {

    for v := range *other {
        (*set)[v] = void{}
    }
}


func (set *Set) Union1(other *Set) {

    diff := other.Diff(set)
    for v := range *diff {
        (*set)[v] = void{}
    }
}


func (set *Set) Union2(other *Set) {

    if set.IsEmpty() {
        *set = *other
    } else if set.Card() > other.Card() {
        for v := range *other {
            if !set.Contains(v) {
                set.Add(v)
            }
        }
    } else {
        for v := range *set {
            if !other.Contains(v) {
                other.Add(v)
            }
        }
        *set = *other
    }
}

func (set *Set) UnionNew(other *Set) (union *Set) {

    union = New()

    for v := range *set {
        union.Add(v)
    }
    for v := range *other {
        union.Add(v)
    }
    return
}

func (set *Set) UnionNew2(other *Set) (union *Set) {

    var diff *Set
    if set.Card() > other.Card() {
        union = set.Clone()
        diff = other.Diff(set)

    } else {
        union = other.Clone()
        diff = set.Diff(other)
    }
    for v := range *diff {
        (*union)[v] = void{}
    }

    return
}


func (set *Set) IsSubset(other *Set) bool {

    if set.Card() > other.Card() {
        return false
    }
    for v := range *set {
        if !other.Contains(v) {
            return false
        }
    }
    return true
}

func (set *Set) IsSuperset(other *Set) bool {
    return other.IsSubset(set)
}

func (set *Set) Clone() (cloned *Set) {
    cloned = New()
    for v := range *set {
        cloned.Add(v)
    }
    return
}

func (set *Set) Clear() {
    *set = *New()
}

func (set *Set) Pop() (item interface{}) {
    for item := range *set {
        delete(*set, item)
        return item
    }
    return nil
}


