package main

import (
    "fmt"
)

type point struct {
    x int
    y int
}

func compare() (bool, bool) {
    point1 := struct {
        x int
        y int
      }{
        10,
        10,
      }
    point2 := struct {
        x int
        y int
    }{}
    point2.x = 10
    point2.y = 5
    
    point3 := point{10, 10}
    
    return point1 == point2, point1 == point3
}

func main() {
    // Title:
    fmt.Println ("Exercise 4.18: Comparing Structs to Each Other")
    fmt.Println("Pag. 152 - PDF:185\n\r")
    // Possible states: Started, Finished
    fmt.Println("Status: Finished\n\r")
    
    a, b := compare()
    fmt.Println("point1 == point2:", a)
    fmt.Println("point1 == point3:", b)

} // End main()
