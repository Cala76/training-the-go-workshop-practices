package main

import (
    "errors"
    "fmt"
)

func doubler(v interface{}) (string, error) {
    if i, ok := v.(int); ok {
        return fmt.Sprint(i * 2), nil
    }
    if s, ok := v.(string); ok {
        return s + s, nil
    }
    return "", errors.New("unsupported type passed")
}

func main() {
    // Title:
    fmt.Println ("Exercise 4.21: Type Assertion")
    fmt.Println("Pag. 161 - PDF:194\n\r")
    // Possible states: Started, Finished
    fmt.Println("Status: Finished\n\r")
    
    res, _ := doubler(5)
    fmt.Println("5   :", res)
    res, _ = doubler("yum")
    fmt.Println("yum :", res)
    _, err := doubler(true)
    fmt.Println("true:", err)

} // End main()
