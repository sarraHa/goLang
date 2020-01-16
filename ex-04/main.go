package main

import (
    "io"
    "os"

    "github.com/margo/rot13"
)

func main() {

    r := rot13.NewReader(os.Stdin)
    io.Copy(os.Stdout, r)

}
