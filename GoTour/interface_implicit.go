package main

import(
	"fmt"
	"math"
)

type I interface{
	M()
}

type T struct{
	S string
}

func (t *T) M(){
	if t==nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type MyFloat float64

func (f MyFloat) M(){
	fmt.Println(f)
}

func main(){
	var i I
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = MyFloat(math.Pi)
	describe(i)
	i.M()
}

func describe(i I){
	fmt.Printf("(%v %T)\n",i,i)
}


