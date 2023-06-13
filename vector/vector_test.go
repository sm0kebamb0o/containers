package vector

import (
	"testing"
)

func TestSort(t *testing.T) {

	cases := []struct {
		in  Vector[int]
		out Vector[int]
	}{
		{Vector[int]{5, 2, 4, 3, 1}, Vector[int]{1, 2, 3, 4, 5}},
		{Vector[int]{3, 4, 2}, Vector[int]{2, 3, 4}},
	}

	for _, c := range cases {
		c.in.Sort(0, c.in.Size(), func(i, j int) bool {
			return c.in[i] < c.in[j]
		})
		for i := range c.in {
			if c.in[i] != c.out[i] {
				t.Errorf("The problem occured on vector %v.", c.in)
				break
			}
		}
	}
}
