package main 
import "fmt"
import "math"


func Uint8FromInt(n int )(uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of uint8 range",n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x  && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x) 
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	} 
	panic(fmt.Sprintf("%g is out of the int32 range",x))
}

func main () {
	var a int = 33
	var b float64 = 11.611
	c, err := Uint8FromInt(a)
	if(err != nil) {
		fmt.Println(err)
	}
	d := IntFromFloat64(b)
	fmt.Printf("%d\n",c)
	fmt.Printf("%d\n",d)

	// 这里引入复数的概念
	// 复数的数据类型是complex32/complex64
	var c1 complex64 = 5 + 5i
	c2 := complex(3.3,33.3)
	// 这里%v作为复数来表示
	fmt.Printf("complex is %v\n",c1)
	fmt.Printf("complex2 is %v\n",c2)
}