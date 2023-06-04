package maps_test

import (
	"testing"

	"github.com/dolanor/gene/maps"
)

func TestClone(t *testing.T) {
	m := map[string]int{
		"1": 1,
		"2": 2,
	}

	clone := maps.Clone(m)
	for k, v := range m {
		got, ok := clone[k]
		if !ok || got != v {
			t.Fatalf("missing data in cloned map: [%s]: %d", k, v)
		}

		want := m[k]
		if &got == &want {
			t.Fatalf("value address are the same. They are not cloned: %p == %p", &got, &want)
		}
	}
}
