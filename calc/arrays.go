package calc

import "math"

func GetCenterOffset(size int, point int) int {
	if point > size {
		return -1
	}
	if size%2 == 0 {
		leftCenterIndex := size/2 - 1
		rightCenterIndex := size / 2

		if point <= leftCenterIndex {
			return leftCenterIndex - point
		} else {
			return point - rightCenterIndex
		}
	}

	centerIndex := size / 2
	return int(math.Abs(float64(centerIndex - point)))
}
