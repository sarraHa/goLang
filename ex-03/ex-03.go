
package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {

	zn := 1.0

	for i := 0; i < 10; i++ {
		
		zn = zn - ( (zn*zn) - x )/2*zn

	}
	
	return zn
}

func main() {
    fmt.Println(Sqrt(2))
}
