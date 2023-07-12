package linearsearch

import "testing"

func TestLinearSearchExistingValue(t *testing.T) {
    haystack := []int{1,2,3,4,5}
    got := LinearSearch(haystack, 3)
    want := true

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }
}

func TestLinearSearchNotExistingValue(t *testing.T) {
    haystack := []int{1,2,3,4,5}
    got := LinearSearch(haystack, 100)
    want := false 

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }
}
