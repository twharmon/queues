package queues_test

import (
	"reflect"
	"testing"

	"github.com/twharmon/queues"
)

func assertEqual(t *testing.T, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestDequeue(t *testing.T) {
	q := queues.New[string]()
	q.Enqueue("foo")
	q.Enqueue("bar")
	assertEqual(t, "foo", q.Dequeue())
	assertEqual(t, "bar", q.Dequeue())
	assertEqual(t, "", q.Dequeue())
}

func TestPeek(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		q := queues.New[string]()
		q.Enqueue("foo")
		q.Enqueue("bar")
		assertEqual(t, "foo", q.Peek())
		assertEqual(t, "foo", q.Peek())
		assertEqual(t, 2, q.Len())
	})
	t.Run("empty", func(t *testing.T) {
		q := queues.New[string]()
		assertEqual(t, "", q.Peek())
	})
}

func TestLen(t *testing.T) {
	q := queues.New[string]()
	q.Enqueue("foo")
	q.Enqueue("bar")
	assertEqual(t, 2, q.Len())
	q.Dequeue()
	assertEqual(t, 1, q.Len())
	q.Dequeue()
	assertEqual(t, 0, q.Len())
}

func TestReslice(t *testing.T) {
	q := queues.New[string]()
	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("baz")
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("baz")
	assertEqual(t, "foo", q.Dequeue())
}

func TestSlice(t *testing.T) {
	q := queues.New[string]()
	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("baz")
	assertEqual(t, []string{"foo", "bar", "baz"}, q.Slice())
}

func TestClear(t *testing.T) {
	q := queues.New[string]()
	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("baz")
	q.Clear()
	assertEqual(t, 0, q.Len())
}

func TestString(t *testing.T) {
	q := queues.New[string]()
	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("baz")
	q.Dequeue()
	assertEqual(t, "[bar baz]", q.String())
}
