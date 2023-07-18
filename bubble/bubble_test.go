package bubble

import "testing"

func TestBubbleSort(t *testing.T) {
    haystack := []int{5,3,4,1,2}
    BubbleSort(haystack)
    want := []int{1,2,3,4,5}

    for i, v := range haystack {
        if v != want[i] {
            t.Errorf("got %v, wanted %v", v, want[i])
        }
    }
}
