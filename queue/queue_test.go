package queue

import "testing"

func TestAddAndPoll(t *testing.T) {
	q := New[int]()

	q.Add(1, 2, 3)

	elem := q.Poll()

	expected := 3

	if elem != expected {
		t.Fatal("Expected", expected, "but got", elem)
	}
}

func TestIsEmpty(t *testing.T) {
	q := New[int]()

	q.Add(1)

	q.Poll()

	if !q.IsEmpty() {
		t.Fatal("Expected queue to be empty")
	}
}

func TestSize(t *testing.T) {
	q := New[int]()

	q.Add(1)

	if q.Size() != 1 {
		t.Fatal("Expected queue to have 1 element")
	}
}

func TestPeek(t *testing.T) {
	q := New[int]()

	q.Add(1)

	elem := q.Peak()

	if elem != 1 {
		t.Fatal("Expected element to be 1 but was", elem)
	}

	if q.Size() != 1 {
		t.Fatal("Expected queue to have 1 element")
	}
}
