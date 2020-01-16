package main

import (
    "fmt"
    "os"
)

// main is function
func main() {
    fmt.Printf("args: %v\n", os.Args)

    // for i := 0; i < len(os.Args); i++ {
    //     fmt.Printf(i, os.Args[i])
    // }

    for i, v := range os.Args[1: ]{
        fmt.Printf("%d: \"%s\"\n", i, v)
    }
}


