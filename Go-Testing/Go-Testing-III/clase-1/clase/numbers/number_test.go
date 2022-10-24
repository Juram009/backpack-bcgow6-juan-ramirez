package numbers

import "testing"

type TestCase struct {
	in  int
	out string
}

// arrange
var tests = []TestCase{{-1, "negative"}, {5, "positive"}}

func TestTypeNumber(t *testing.T) {
	for i, test := range tests {
		// act
		size := TypeNumber(test.in)
		// assert
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
