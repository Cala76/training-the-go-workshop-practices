package main

import (
    "fmt"
    "math"
)

func convert() string{
    var i8 int8 = math.MaxInt8
    i := 128
    f64 := 3.14
     m := fmt.Sprintf("int8  = %v > int64  = %v\n", i8, int64(i8))
     // This will cause an overflow to int8's minimum size:
     m += fmt.Sprintf("int   = %v > int8   = %v\n", i, int8(i))
     m += fmt.Sprintf("int8  = %v > float64 = %v\n", i8, float64(i8))
     m += fmt.Sprintf("float64 = %v > int   = %v\n", f64, int(f64))
     return m
}

func main() {
    // Title:
    fmt.Println ("Exercise 4.20: Numeric Type Conversion")
    fmt.Println("Pag. 158 - PDF:191\n\r")
    // Possible states: Started, Finished
    fmt.Println("Status: Finished\n\r")
    
    fmt.Print(convert())
  
} // End main()
