package main

import "testing"

// func TestSqrt(t *testing.T) {
// 	const in, out = 4, 2
// 	if x := Sqrt(in); x != out {
// 		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
// 	}
// }

func TestTwo(t *testing.T) {
	const in, out = 100, 2
	if x := Two(in); x != out {
		t.Errorf("Two(%v) = %v, want %v", in, x, out)
	}

	const in, out = 550000, 2
	if x := Two(in); x != out {
		t.Errorf("Two(%v) = %v, want %v", in, x, out)
	}

	const in, out = 111, 2
	if x := Two(in); x != out {
		t.Errorf("Two(%v) = %v, want %v", in, x, out)
	}
}