package main

import (
	"fmt"
)

type human struct {
	sex             string
	age, boneLength float64
}

func (h human) height() float64 {
	if h.sex == "m" {
		return h.boneLength*2.238 + 69.089
	}

	return h.boneLength*2.317 + 61.412
}

func main() {
	var (
		age, boneLength float64
		sex             string
	)

	fmt.Print("Enter sex: ")
	_, err := fmt.Scanf("%s", &sex)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter your age: ")
	_, err = fmt.Scanf("%f", &age)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter bone lenght: ")
	_, err = fmt.Scanf("%f", &boneLength)
	if err != nil {
		panic(err)
	}

	h := human{
		age:        age,
		boneLength: boneLength,
		sex:        sex}

	fmt.Println("Your height is ", h.height(), "cm")

}
