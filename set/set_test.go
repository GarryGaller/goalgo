package set

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	got := New()

	if got == nil || got == new(Set) {
		t.Errorf("set.New() = %v; want %v", got, New())
	}
}

func TestClone(t *testing.T) {

	set := New(1)
	got := set.Clone()

	if !reflect.DeepEqual(set, got) {
		t.Errorf("set.Clone() = %v; want %v", got, set)
	}

}

func TestClear(t *testing.T) {

	set := New(1)
	set.Clear()

	if !reflect.DeepEqual(set, New()) {
		t.Errorf("set.Clear() = %#v; want %#v", set, New())
	}
}

func TestAdd(t *testing.T) {

	set := New()
	set.Add(1)

	if !reflect.DeepEqual(set, New(1)) {
		t.Errorf("set.Add(1) = %v; want %v", set, New(1))
	}

}

func TestPop(t *testing.T) {

	set := New(1)
	set.Pop()

	if !set.IsEmpty() {
		t.Errorf("set.Pop() = %v; want %v", set, New())
	}

}

func TestContains(t *testing.T) {

	set := New(1)
	got := set.Contains(1)
	if !got {
		t.Errorf("set.Contains(1) = %v; want %v", got, true)
	}

}

func TestDelete(t *testing.T) {
	set := New(1)
	set.Delete(1)

	if !set.IsEmpty() {
		t.Errorf("set.Delete(1) = %v; want %v", set, New())
	}
}

func TestInter(t *testing.T) {

	set := New(1, 2, 3)
	other := New(1, 2)

	got := set.Inter(other)

	if !reflect.DeepEqual(New(1, 2), got) {
		t.Errorf("set.Inter() = %v; want %v", got, New(1, 2))
	}
}

func TestDiff(t *testing.T) {

	set := New(1, 2, 3)
	other := New(1, 2)

	got := set.Diff(other)

	if !reflect.DeepEqual(New(3), got) {
		t.Errorf("set.Diff() = %v; want %v", got, New(3))
	}
}

func TestSymDiff(t *testing.T) {

	set := New(1, 2, 3)
	other := New(0, 1, 2)

	got := set.SymDiff(other)

	if !reflect.DeepEqual(New(0, 3), got) {
		t.Errorf("set.SymDiff() = %v; want %v", got, New(0, 3))
	}
}

func TestUnion(t *testing.T) {

	set := New(1, 2, 3)
	other := New(4)

	set.Union(other)

	if !reflect.DeepEqual(New(1, 2, 3, 4), set) {
		t.Errorf("set.SymDiff() = %v; want %v", set, New(1, 2, 3, 4))
	}
}

func UnionNew(t *testing.T) {

	set := New(1, 2, 3)
	other := New(4)

	got := set.UnionNew(other)

	if !reflect.DeepEqual(New(1, 2, 3, 4), set) {
		t.Errorf("set.SymDiff() = %v; want %v", got, New(1, 2, 3, 4))
	}
}

func TestIsSubset(t *testing.T) {

	set := New(1, 2)
	other := New(1, 2, 3)

	got := set.IsSubset(other)

	if !got {
		t.Errorf("set.IsSubset() = %v; want %v", got, true)
	}
}

func TestIsSuperset(t *testing.T) {

	set := New(1, 2, 3)
	other := New(1, 2)

	got := set.IsSuperset(other)

	if !got {
		t.Errorf("set.IsSuperset() = %v; want %v", got, true)
	}
}

func TestSize(t *testing.T) {

	set := New(1)
	got := set.Size()

	if got != 1 {
		t.Errorf("set.Size() = %v; want %v", got, 1)
	}

}

func TestCard(t *testing.T) {

	set := New(1)
	got := set.Card()

	if got != 1 {
		t.Errorf("set.Card() = %v; want %v", got, 1)
	}
}

func TestIsEmpty(t *testing.T) {

	t.Run("Empty", func(t *testing.T) {
		set := New()
		got := set.IsEmpty()
		if !got {
			t.Errorf("set.IsEmpty() = %v; want %v", got, true)
		}
	})

	t.Run("Not Empty", func(t *testing.T) {
		set := New(1)
		got := set.IsEmpty()
		if got {
			t.Errorf("set.IsEmpty() = %v; want %v", got, false)
		}
	})

}
