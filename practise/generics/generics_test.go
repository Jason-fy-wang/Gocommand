package generics

import (
	"testing"
)

func Test_Print(t *testing.T) {
	Print(1)
	Print("hello")
}

func Test_SumMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	sum := SumMap(m)
	if sum != 6 {
		t.Errorf("expected 6, got %d", sum)
	}
}

func Test_Add(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Errorf("expected 3, got %d", sum)
	}
}

func Test_AddNumber(t *testing.T) {
	sum := AddNumber(1, 3)

	if sum != 4 {
		t.Errorf("expected 4, got %d", sum)
	}
}

func Test_Container(t *testing.T) {
	c := NewContainer(10)
	if c.Value != 10 {
		t.Errorf("expected 10, got %d", c.Value)
	}

	c.SetVal(50)
	if c.Value != 50 {
		t.Errorf("expected 50, got %d", c.Value)
	}

}

func Test_Map(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	res := Map(arr, func(v int) int {
		v *= 2
		return v
	})

	if len(res) != len(arr) {
		t.Errorf("expected length %d, got %d", len(arr), len(res))
	}

	for i, v := range res {
		if v != arr[i]*2 {
			t.Errorf("expected %d, got %d", arr[i]*2, v)
		}
	}

}
