package db

import "testing"

func TestStringStore(t *testing.T) {
	store := newStringStore()

	expectations := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}

	store.set("k1", expectations["k1"])
	store.set("k2", expectations["k2"])
	store.set("k3", expectations["k3"])

	for ek, ev := range expectations {
		actual := store.get(ek)
		if actual != ev {
			t.Fatalf(`Actual value: %s for %s, expected: %s`, actual, ek, expectations[ek])
		}
	}
}
