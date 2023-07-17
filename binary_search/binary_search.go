package binarysearch

func BinarySearch(haystack []int, needle int) bool {
    low, high := 0, len(haystack) - 1

    for low < high {
        mid := low + (high - low)/2

        if haystack[mid] == needle {
            return true
        } else if haystack[mid] > needle {
            high = mid 
        } else {
            low = mid + 1
        }
    }

    return false
}
