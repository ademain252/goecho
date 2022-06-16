package basic

// Slice generic type 泛型类型
// T: type parameter
// int | float32 | float64: type constraint
type Slice[T int | float32 | float64] []T

// Instantiations 泛型类型实例化
var weekday Slice[int] = []int{1, 2, 3, 4, 5}

func (s Slice[T]) Sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
