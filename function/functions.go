package function

import "fmt"

func Average(xs []float64) (avg float64) {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	avg = total / float64(len(xs))
	return
}

//Variadic functions
func Add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//Closure
func MakeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func Factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * Factorial(x-1)
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func DeferFunction() {
	defer second()
	first()
}

func PanicFunction() {
	defer func() {
		str := recover()
		fmt.Printf("Recovered value is %s\n", str)
	}()
	panic("PANIC")
}

func zero(xPtr *int) {
	*xPtr = 0
}

func PointerTest() {
	x := 5
	zero(&x)
	fmt.Println(x)
}

func one(xPtr *int) {
	*xPtr = 1
}

func NewFunctionTest() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}

func Half(input int) (value int, isEven bool) {
	value = input / 2
	isEven = input%2 == 0
	return
}

func Square(x *float64) {
	*x = *x * *x
}

func Swap(x *int, y *int) {
	tmp := *y
	*y = *x
	*x = tmp
}
