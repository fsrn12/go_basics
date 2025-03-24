package lib

func Add(num1 int, num2 int) int{
	 result := num1 + num2
	 return result
}

func CalculateTotal(item1, item2 float64)float64 {
	return float64(item1) + float64(item2)
}

// multiple return types from a function
// every return value is a first-class citizen
func CalculateSumAndProduct(item1, item2 int) (int, int){
	return item1 + item2, item1*item2
}
