package main

import (
	"fmt"
	"testing"
)

func TestAllEvenBuggy(t *testing.T) {
	testCases := []int{1, 2, 4, 6}
	for _, v := range testCases {
		t.Run("sub", func(t *testing.T) {
			t.Parallel()
			t.Logf("testing v=%d", v)
			if v&1 != 0 {
				t.Fatal("odd v", v)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	var f func()
	for i := 0; i < 10; i++ {
		if i == 0 {
			f = func() { fmt.Println(i) }
		}
		fmt.Println(i)
		f()
	}
}
