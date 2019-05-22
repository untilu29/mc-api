package hello_world

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func split(sum int) (x, y int) {
	x = (sum - 2) * 3
	y = sum - x
	return
}

var chuc, mi, diem, viet string

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-6 + 9i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println("Hello, world")
	fmt.Println("Time", time.Now())
	rand.Seed(4)
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println("Cong 2 so", add(6, 9))
	a, b := swap(4, 6)
	fmt.Println(a, b)
	fmt.Println(split(45))
	fmt.Println(chuc, mi, diem, viet)

	var x, y, c, d int = 1, 2, 3, 5

	fmt.Println(x, y, c, d)

	boolData := "Chuc"
	fmt.Println(boolData)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var l, m int = 3, 4
	var f float64 = math.Sqrt(float64(l*l + m*m))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	fmt.Println("----------------------------")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println("--------For-condition-------")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum ", sum)

	for sum < 100 {
		sum += 1
	}
	fmt.Println("New sum", sum)

	//for {
	//
	//}        ------- Forever loop

	//-----------------------------Condition--------------------
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(2, 5, 15))

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
