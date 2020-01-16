

package main

import (
    "flag"
    "fmt"
	"golang.org/x/tour/tree"
)

type Point struct{
	x, y int
}

var (
	mu syn

)
func main() {
    nn := flag.Int("xy", 2, "NxN size of the grid")
    flag.Parse()

    nx := *nn
    ny := *nn

    fmt.Printf("n=%d\n", compute(nx, ny))
}


func inspect(pt Point, nx, ny int){



}
func compute(nx, ny int) int {
    
	var pt Point
	r := Point{ pt.x+1, pt.y}
	r := Point{ pt.x=, pt.y+1}

	go inspect(r, nx, ny )
	go inspect(d, nx, ny )

	ch := make(chan Point) 
	t := tree.Tree{
		
	}	
	

    return 0
}


