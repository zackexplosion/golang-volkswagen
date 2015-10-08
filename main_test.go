package main

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

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
}

func TestSpec(t *testing.T) {

    // Only pass t into top-level Convey calls
    Convey("Given some integer with a starting value", t, func() {
        x := 1

        Convey("When the integer is incremented", func() {
            x++

            Convey("The value should be greater by one", func() {
                So(x, ShouldEqual, 2)
            })
        })
    })
}