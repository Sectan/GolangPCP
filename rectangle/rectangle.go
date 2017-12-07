package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type geometry interface {
	Area() float64
	Perim() float64
}

type rect struct {
	width, height float64
}

func (r *rect) Area() float64 {
	return r.width * r.height
}
func (r *rect) Perim() float64 {
	return 2*r.width + 2*r.height
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("no arguments")
		os.Exit(1)
	}

	var methodName string
	methodName = os.Args[1]

	w, e := strconv.ParseFloat(os.Args[2], 64)
	if e != nil {
		fmt.Println("Error parsing width")
		panic(e)
	}

	h, e := strconv.ParseFloat(os.Args[3], 64)
	if e != nil {
		fmt.Println("Error parsing height")
		panic(e)
	}

	r := rect{width: w, height: h}
	//fmt.Println(r.Area())
	v := reflect.ValueOf(&r).MethodByName(methodName).Call([]reflect.Value{})
	fmt.Println(v[0].Float())
}
