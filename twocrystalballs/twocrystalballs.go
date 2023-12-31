package twocrystalballs

import "math"

func TwoCrystalBalls(breaks []bool) int {
    size := float64(len(breaks))
    jumpSize := int(math.Floor(math.Sqrt(size)))

    i := jumpSize

    for i < len(breaks) {
        if (breaks[i]) {
            break
        }
        i += jumpSize
    }

    i -= jumpSize

    for j := 0; j < jumpSize && i <= len(breaks); i,j = i+1, j+1 {
        if (breaks[i]) {
            return i;
        }
    }

    return -1
}
