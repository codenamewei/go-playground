package main

import (
    "fmt"
    "github.com/codenamewei/go-playground/hello-world/build-custom-package"
)

func main() {
    result := mypackage.Add(3, 4)
    fmt.Println("Result:", result)
}