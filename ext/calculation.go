package ext

func CalcCost(price float32, current_value float32, previous_value float32) float32 {
	return price * (current_value - previous_value)
}
