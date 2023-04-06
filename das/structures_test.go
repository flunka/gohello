package das

import (
	"fmt"
	"testing"
)

func TestQueuePrint(t *testing.T) {
	cases := []struct {
		in   Queue[int]
		want string
	}{
		{Queue[int]{head: &QueueNode[int]{value: 1, next: &QueueNode[int]{value: 2}}}, "1->2"},
		{Queue[int]{head: &QueueNode[int]{value: 1}}, "1"},
	}
	for _, c := range cases {
		got := fmt.Sprint(c.in)
		if got != c.want {
			t.Errorf("Got %s want %s", got, c.want)
		}
	}
}

func TestEnqueue(t *testing.T) {
	queue := Queue[int]{head: nil, tail: nil}
	queue.enqueue(1)
	queue.enqueue(2)
	want := "1->2"
	got := fmt.Sprint(queue)
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
	queue.enqueue(3)
	queue.enqueue(4)
	want = "1->2->3->4"
	got = fmt.Sprint(queue)
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func TestNewQueue(t *testing.T) {
	cases := []struct {
		in   []int
		want string
	}{
		{[]int{1, 2}, "1->2"},
		{[]int{1, 2, 3, 4}, "1->2->3->4"},
		{[]int{}, ""},
	}
	for _, c := range cases {
		queue := NewQueue[int](c.in)
		got := fmt.Sprint(queue)
		if got != c.want {
			t.Errorf("Got %s want %s", got, c.want)
		}
	}
}

func TestDeque(t *testing.T) {
	queue := NewQueue[int]([]int{1, 2, 3, 4})
	got_value, _ := queue.deque()
	want_value := 1
	if got_value != want_value {
		t.Errorf("Got %d want %d", got_value, want_value)
	}
	want := "2->3->4"
	got := fmt.Sprint(queue)
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
	queue.deque()
	queue.deque()
	got_value, _ = queue.deque()
	want_value = 4
	if got_value != want_value {
		t.Errorf("Got %d want %d", got_value, want_value)
	}
	want = ""
	got = fmt.Sprint(queue)
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
	_, err := queue.deque()
	want_error := ErrEmptyQueue
	if err == nil {
		t.Error("wanted an error but didn't get one")
	} else if err != want_error {
		t.Errorf("Got: %s want: %s", err, want_error)
	}

}

func TestStackPrint(t *testing.T) {
	cases := []struct {
		in   Stack[int]
		want string
	}{
		{Stack[int]{top: &StackNode[int]{value: 1, prev: &StackNode[int]{value: 2}}}, "1->2"},
		{Stack[int]{top: &StackNode[int]{value: 1}}, "1"},
		{Stack[int]{top: nil}, ""},
	}
	for _, c := range cases {
		got := fmt.Sprint(c.in)
		if got != c.want {
			t.Errorf("Got %s want %s", got, c.want)
		}
	}
}

func TestPush(t *testing.T) {
	stack := Stack[int]{top: nil}
	stack.push(1)
	stack.push(2)
	want := "2->1"
	got := fmt.Sprint(stack)
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
	stack.push(3)
	stack.push(4)
	want = "4->3->2->1"
	got = fmt.Sprint(stack)
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func TestPop(t *testing.T) {
	stack := Stack[int]{top: nil}
	stack.push(1)
	stack.push(2)
	want := "2->1"
	got := fmt.Sprint(stack)
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
	got_value, _ := stack.pop()
	want_value := 2
	if got_value != want_value {
		t.Errorf("Got %d want %d", got_value, want_value)
	}
	stack.pop()
	want = ""
	got = fmt.Sprint(stack)
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
	_, err := stack.pop()
	want_error := ErrEmptyStack
	if err == nil {
		t.Error("wanted an error but didn't get one")
	} else if err != want_error {
		t.Errorf("Got %s want %s", err, want_error)
	}
}
