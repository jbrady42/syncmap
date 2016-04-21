package syncmap

import "testing"

func TestNew(t *testing.T) {
	m := New()
	if m == nil {
		t.Error("Map is nil")
	}
}

func TestSet(t *testing.T) {
	m := New()
	m.Set("test", 1)

	item, ok := m.data["test"]

	if !ok {
		t.Error("Data should contain item")
	}

	if item != 1 {
		t.Error("test should = 1")
	}
}

func TestHas(t *testing.T) {
	m := New()
	m.Set("test", 1)

	if !m.Has("test") {
		t.Error("Data should contain item")
	}
}

func TestGet(t *testing.T) {
	m := New()
	m.Set("test", 1)

	item, ok := m.Get("test")

	if !ok {
		t.Error("Data should contain item")
	}

	if item != 1 {
		t.Error("test should = 1")
	}
}

func TestLen(t *testing.T) {
	m := New()
	m.Set("test", 1)
	m.Set("test2", 2)

	count := m.Len()

	if count != 2 {
		t.Error("Count should be 2")
	}
}

func TestDelete(t *testing.T) {
	m := New()
	m.Set("test", 1)
	m.Set("test2", 2)

	m.Delete("test2")

	if m.Has("test2") {
		t.Error("Key should have been removed")
	}
}

func TestIter(t *testing.T) {
	m := New()
	m.Set("test", 1)
	m.Set("test2", 2)

	count := 0
	for tup := range m.Iter() {
		count += tup.value.(int)
	}

	if count != 3 {
		t.Error("Should iterate all items")
	}
}
