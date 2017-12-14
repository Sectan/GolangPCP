package main

import "fmt"
import "github.com/Sectan/GolangPCP/mymath"

type human struct {
	sex             string
	age, boneLength float64
}

func (h human) height() float64 {
	if h.sex == "m" {
		return mymath.Addition(mymath.Multiply(h.boneLength,2.238), 69.089)
	}

	return mymath.Addition(mymath.Multiply(h.boneLength,2.317), 61.412)
}

func main() {
	var (
		age, boneLength float64
		sex             string
	)

	fmt.Print("Enter sex: ")
	_, err := fmt.Scanf("%s\n", &sex)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter your age: ")
	_, err = fmt.Scanf("%f\n", &age)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter bone lenght: ")
	_, err = fmt.Scanf("%f\n", &boneLength)
	if err != nil {
		panic(err)
	}

	h := human{
		age:        age,
		boneLength: boneLength,
		sex:        sex}

	fmt.Println("Your height is ", h.height(), "cm")

}
