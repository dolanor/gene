package slices_test

import (
	"testing"

	"github.com/dolanor/gene/slices"
)

func TestClone(t *testing.T) {
	m := []string{
		"0",
		"1",
		"2",
		"3",
	}

	clone := slices.Clone(m)
	for k, v := range m {
		got := clone[k]
		if got != v {
			t.Fatalf("missing data in cloned map: %s", v)
		}

		want := m[k]
		if &got == &want {
			t.Fatalf("value address are the same. They are not cloned: %p == %p", &got, &want)
		}
	}
}

func BenchmarkClone(b *testing.B) {
	benchClone(b, slices.Clone[string])
}

func BenchmarkCloneWithSize(b *testing.B) {
	benchClone(b, cloneWithSize[string])
}

func BenchmarkCloneWithoutMake(b *testing.B) {
	benchClone(b, cloneWithoutMake[string])
}

func benchClone(b *testing.B, f func([]string) []string) {
	tt := make([]string, 1000)

	for i := 0; i < 1000; i++ {
		tt[i] = "champagne"
	}

	for i := 0; i < b.N; i++ {
		f(tt)
	}
}

func cloneWithSize[T comparable](tt []T) []T {
	clone := make([]T, len(tt))
	for k, v := range tt {
		clone[k] = v
	}
	return clone
}

func cloneWithoutMake[T comparable](s []T) []T {
	clone := []T{}
	for _, v := range s {
		clone = append(clone, v)
	}
	return clone
}
