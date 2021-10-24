package main

// a simple sqrt implementation 
import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z:= 1.0;
	prevz := 1.0 - 0.01;
	var i uint64;
	for ;math.Abs(z - prevz) > 0.001; i++ {
		prevz = z;
		z -= (z*z - x) / (2*z);
	}
	fmt.Printf("took %d iterations\n", i);
	return z;
}

func main() {
	fmt.Println(Sqrt(2));
}
