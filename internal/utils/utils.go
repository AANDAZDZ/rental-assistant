package utils

import "math"

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Haversine(lat1, lon1, lat2, lon2 float64) int {
	lat1, lon1, lat2, lon2 = toRadians(lat1), toRadians(lon1), toRadians(lat2), toRadians(lon2)
	tmpLat := lat2 - lat1
	tmpLon := lon2 - lon1
	a := math.Sin(tmpLat/2)*math.Sin(tmpLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(tmpLon/2)*math.Sin(tmpLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	r := 6378.137
	distance := r * c * 1000
	return int(distance)
}
