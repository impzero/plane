package calc

import "math"

func CenterOffset(arraySize int, arrayPoint int) int {
	if arrayPoint > arraySize {
		return -1
	}
	if arraySize%2 == 0 {
		leftCenterIndex := arraySize/2 - 1
		rightCenterIndex := arraySize / 2

		if arrayPoint <= leftCenterIndex {
			return leftCenterIndex - arrayPoint
		} else {
			return arrayPoint - rightCenterIndex
		}
	}

	centerIndex := arraySize / 2
	return int(math.Abs(float64(centerIndex - arrayPoint)))
}
