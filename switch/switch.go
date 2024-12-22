package main

import (
    "fmt"
    "time"
)

func main() {

    i := 2
    fmt.Print("Write ", i, " as ")
    switch i { //basic switch statement
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() { 
    case time.Saturday, time.Sunday: //multiple statements comma separated
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12: //if else using switch
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    //WhatAmI function declaration
    whatAmI := func(i interface{}) { //interface value will correspond to clause and explained later
        switch t := i.(type) { //type switch
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
