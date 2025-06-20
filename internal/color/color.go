package color

import "math"

func closeEnough(f float64) float64 {
	if math.Abs(math.Round(f)-f) < 1e-5 {
		return math.Round(f)
	}
	return f
}
