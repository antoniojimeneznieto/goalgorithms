package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kWh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var intSlice []int = make([]int, 0, 8)
	intSlice = append(intSlice, 1, 2, 3, 4, 5)

	var float32Slice []float32 = make([]float32, 0, 8)
	float32Slice = append(float32Slice, 1.1, 2.2, 3.3, 4.4, 5.5)

	var float64Slice []float64 = make([]float64, 0, 8)
	float64Slice = append(float64Slice, 1.1, 2.2, 3.3, 4.4, 5.5)

	var intSum = sumSlice(intSlice)
	var float32Sum = sumSlice(float32Slice)
	var float64Sum = sumSlice(float64Slice)

	fmt.Printf("Sum of intSlice: %d\n", intSum)
	fmt.Printf("Sum of float32Slice: %f\n", float32Sum)
	fmt.Printf("Sum of float64Slice: %f\n", float64Sum)

	var gasCar = car[gasEngine]{
		carMake:  "Toyota",
		carModel: "Corolla",
		engine: gasEngine{
			gallons: 10,
			mpg:     30,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model S",
		engine: electricEngine{
			kWh:   60,
			mpkwh: 3.5,
		},
	}

	fmt.Printf("Gas car: %v\n", gasCar)
	fmt.Printf("Electric car: %v\n", electricCar)
}

func sumSlice[T int | float32 | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}

	return sum
}

func isEmpty[T any](s []T) bool {
	return len(s) == 0
}
