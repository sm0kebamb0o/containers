package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	cases := []struct {
		in  Stack[int]
		out Stack[int]
	}{
		{Stack[int]{1, 2, 3}, Stack[int]{1}},
		{Stack[int]{1, 2}, Stack[int]{}},
	}

	for j, c := range cases {
		for i, n := 0, c.in.Size()-c.out.Size(); i < n; i++ {
			c.in.Pop()
		}

		for i := 0; i < c.in.Size(); i++ {
			if c.in[i] != c.out[i] {
				t.Errorf("The problem occured on test #%v", j)
			}
		}
	}
}
