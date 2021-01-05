package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.0f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} 
	var iter int
	z := x
	for i := 0; i < 10; i++ {
		iter = i
		v := z
		z -= (z*z - x) / (2*z)
		if v == 1 {
			fmt.Println(z)
			continue
		} else if v - z < 1e-8 {
			break
		}
		fmt.Println(z)
	}
	fmt.Println("Liczba iteracji to", iter)
	return z, nil
}

func main() {
	Sqrt(2)
	fmt.Println(Sqrt(-2))
}