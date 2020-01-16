package main

import (
    "fmt"
    "os"
	"strconv" 
)


// main is function
func main() {
    
	
	sum := 0
	for i := 1; i < len(os.Args); i++ {

		v , err := sum +  strconv.Atoi(os.Args[i])
		
		if err != nil {
      		// handle error
			fmt.Printf("error : %v\n", err)
			continue   		
		}
     }
   
	fmt.Printf("la somme est : ", sum)
}



