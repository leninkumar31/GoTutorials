package main 

import "fmt"
import "math"

type MyFloat float64

func (val MyFloat) Abs() float64{
	if val<0 {
		return float64(-val)
	}
	return float64(val)
}

func main(){
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}