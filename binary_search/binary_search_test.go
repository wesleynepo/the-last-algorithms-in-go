package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
    values := []int{1,3,4,69, 71, 81, 90, 99, 420, 1337, 69420} 

    got := BinarySearch(values, 69)
    want := true

    if got != want {
        t.Errorf("got %v wanted %v", got, want)
    }
}
