package main

func main() {
    arr := []int{9,3,7,4,69, 420}

    quicksort(arr)


    for _, a := range arr {
        println(a)
    }
}

func quicksort(arr []int) {
    qs(arr, 0, len(arr) -1)
}

func qs(arr []int, low, hi int) {
    if (low >= hi) {
        return
    }

    pivot := partition(arr, low, hi)

    qs(arr, low, pivot -1)
    qs(arr, pivot + 1, hi)
}

func partition(arr []int, low, hi int) int {
    pivot := arr[hi

    index := low -1

    for i := low; i < hi; i++ {
        if (arr[i] <= pivot) {
            index++
            arr[i], arr[index] = arr[index], arr[i]
        }
    }

    index++
    arr[hi], arr[index] = arr[index], pivot
    return index
}
