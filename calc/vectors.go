package calc

import "math"

func CalculateMeanDirection(vectors [][2]float64) float64 {
	sumX := 0.0
	sumY := 0.0
	for _, vector := range vectors {
		sumX += vector[0]
		sumY += vector[1]
	}

	// Calculate the mean direction
	meanDirection := math.Atan2(sumY/float64(len(vectors)), sumX/float64(len(vectors)))
	// Convert the angle from radians to degrees
	meanDirectionDegrees := meanDirection * 180 / math.Pi
	// Ensure the angle is between 0 and 360 degrees
	meanDirectionDegrees = math.Mod(meanDirectionDegrees+360, 360)
	return meanDirectionDegrees
}

func GetVector(vectors [][2]float64) [2]float64 {
	sumX := 0.0
	sumY := 0.0

	for _, vector := range vectors {
		sumX += vector[0]
		sumY += vector[1]
	}

	return [2]float64{sumX, sumY}
}
