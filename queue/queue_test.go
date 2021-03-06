package queue

import (
	"testing"
	"fmt"
)

func TestQueue(t *testing.T) {
	q := NewQueue(6)

	for i := 0; i < 35; i++ {
		q.Put(fmt.Sprintf("%d",i))
	}

	for i := 0; i < 35; i++ {
		if q.Get() != fmt.Sprintf("%d",i) {
			t.Fatal("get element err")
		}
	}

	if q.Get() != nil {
		t.Fatal("get unwanted element")
	}
}