package main

import "fmt"

func main() {
    var f1 float32 = 1677.7216 // IEEE-754 32-bit floating-point numbers
    fmt.Println(f1)
    
    var f2 float64 = 18787677.878716 // IEEE-754 64-bit floating-point numbers
    fmt.Println(f2)
    
    var f3 complex64 = 8789579877.721 // Complex numbers with float32 real and imaginary parts
    fmt.Println(f3)
    
    var f4 complex128 = 985214745.7216 // Complex numbers with float64 real and imaginary parts
    fmt.Println(f4)
    
}