package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {
	intro()
	dataTypes()
	loops()
	fmt.Println("Sum Function call:", sum(1, 2))
	returnTypes()
	receivers()
	typeAssertions()
	typeSwitches()
}

func intro() {
	const challenge = "#90DaysOfDevOps!"
	const daysComplete = 11
	const daysTotal = 90

	fmt.Println("Intro: Welcome to", challenge, "This is a", daysTotal, "challenge and you have completed", daysComplete, "days.")
}

func dataTypes() {
	fmt.Println("--Basic Types--")
	integers()

	fmt.Println("--Aggregate Types--")
	arrays()
	structs()

	fmt.Println("--Reference Types--")
	slices()
	maps()
	pointers()
}

func integers() {
	var i int
	var x = 5

	fmt.Println("Integers: ", i, x)
}

func arrays() {
	var emptyArray [5]int

	initializedArray := [5]int{0, 1, 2, 3, 4}
	var firstElement = initializedArray[0]

	fmt.Println("Arrays:", emptyArray, initializedArray, firstElement)
}

func slices() {
	var slice = append([]int{}, 12)

	fmt.Println("Slices:", slice)
}

func maps() {
	var emptyMap = make(map[int]int)
	vertices := make(map[string]int)
	square := vertices["square"]
	noValidkey := vertices["notInTheMap"]

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "triangle")

	fmt.Println("Maps:", emptyMap, vertices, square, "NotInTheMapReturns:", noValidkey)
}

func loops() {
	arr := []string{"a", "b", "c"}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	forLoop := ""
	for i := 0; i < 5; i += 1 {
		forLoop += strconv.FormatInt(int64(i), 10)
	}

	arrayLoop := ""
	for index, value := range arr {
		arrayLoop += strconv.FormatInt(int64(index), 10)
		arrayLoop += ","
		arrayLoop += value
		arrayLoop += " "
	}

	mapLoop := ""
	for key, value := range m {
		mapLoop += key
		mapLoop += value
	}

	fmt.Println("Loops:", forLoop, arrayLoop, mapLoop)
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func returnTypes() {
	values := [2]float64{-16, 16}
	logOutput := "Errors:"

	for _, value := range values {
		result, err := sqrt(value)

		if err != nil {
			logOutput += " Error:"
			logOutput += err.Error()
		} else {
			logOutput += " Result:"
			logOutput += strconv.FormatFloat(result, 'E', -1, 32)
		}
	}

	fmt.Println(logOutput)
}

type Person struct {
	name string
	age  int
}

func structs() {
	p := Person{name: "Thomas", age: 27}
	fmt.Println("Struct of a Person:", p)
}

func pointers() {
	i := 7
	inc(&i)
	fmt.Println("Increased Value by using a pointer:", i)
}

func inc(value *int) {
	*value++
}

type Square struct {
	Width, Height int
}

func (s Square) Area() int {
	return s.Height * s.Width
}

func (s *Square) Scale(k int) {
	s.Width *= k
	s.Height *= k
}

func receivers() {
	square := Square{Width: 2, Height: 5}
	fmt.Println("Receiver test on square:", square)
	fmt.Println("Receiver usage to calc square area:", square.Area())
	square.Scale(2)
	fmt.Println("Receiver usage to calc square area after scaling Height and Width by 2:", square.Area())
}

func typeAssertions() {
	var i interface{} = 20

	s := i.(int)
	fmt.Println("Type Assertions that s is an int:", s)

	s, ok := i.(int)
	fmt.Println("Type Assertions that s is an int with ok:", s, ok)

	f, ok := i.(float64)
	fmt.Println("Type Assertions that s is a float with ok:", f, ok)

	fmt.Println("Type Assertions that s is a float without ok will panic:")
	// f = i.(float64) -> will panic
}

func typeSwitches() {
	elaborateType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("Type Switch case int")
		case string:
			fmt.Println("Type Switch case string")
		default:
			fmt.Println("Type Switch case default:", v)
		}
	}

	elaborateType(10)      // int
	elaborateType("Hello") // string
	elaborateType(64.0)    // float64
}
