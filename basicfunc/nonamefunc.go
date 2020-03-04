package basicfunc

//Squares  匿名函数
func Squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
