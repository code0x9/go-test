package main

import (
	"cmp"
	"fmt"
	"slices"
	"sync"
	"testing"

	"golang.org/x/exp/constraints"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func GMin[T cmp.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

func TestGeneric(t *testing.T) {
	a, b := 3, 5
	expected := 3

	got := Min(a, b)
	if got != expected {
		t.Errorf("Min(%d, %d) = %d; want %d", a, b, got, expected)
	}

	got = GMin(a, b)
	if got != expected {
		t.Errorf("Min(%d, %d) = %d; want %d", a, b, got, expected)
	}

	x, y := 2.5, 1.5
	expectedF := 1.5

	gotF := GMin(x, y)
	if gotF != expectedF {
		t.Errorf("GMin(%f, %f) = %f; want %f", x, y, gotF, expectedF)
	}

	s, c := []int{1, 2, 3}, 3
	expectedS := []int{3, 6, 9}

	gotS := Scale(s, c)
	if slices.Compare(gotS, expectedS) != 0 {
		t.Errorf("Scale(%v, %d) = %v; want %v", s, c, gotS, expectedS)
	}
}

func TestSync(t *testing.T) {
	var a string
	var once sync.Once
	var wg sync.WaitGroup

	setup := func() {
		fmt.Println("setup")
		a = "hello, world..."
	}

	doPrint := func() {
		once.Do(setup)
		fmt.Println(a)
		wg.Done()
	}

	go doPrint()
	go doPrint()
	go doPrint()
	wg.Add(3)

	wg.Wait()

	if len(a) == 0 {
		t.Errorf("a = %s; want %s", a, "hello, world...")
	}
}
