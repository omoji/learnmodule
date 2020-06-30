package simpleinterest

//Calculates the simple interest given the principal (p), the rate of interest (r) and the duration of savings (t)
func Calculate(p, r, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
