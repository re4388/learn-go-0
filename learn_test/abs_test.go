package learn_test

import (
	"hello/learn"
	"testing"
)

func TestCircle(t *testing.T) {
	res := learn.Circle{Radius: 5.0}

	if res.Area() != 78.50 {
		t.Errorf("res.Area() = %0.2f; not 78.50", res.Area())
	}
}
