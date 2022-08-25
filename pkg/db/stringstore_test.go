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

	for k, v := range expectations {
		actual := store.get(k)
		if actual != v {
			t.Fatalf(`Actual value: %s for %s, expected: %s`, actual, k, expectations[k])
		}
	}
}
