package main

import "fmt"


//For is the only looping in Go
func main() {

    i := 1
    for i <= 3 { //single condition
        fmt.Println(i)
        i = i + 1
    }

    for j := 0; j < 3; j++ { //c style
        fmt.Println(j)
    }

    for i := range 3 { //do 3 times
        fmt.Println("range", i) //prints 0,1,2
    }
 
    for { //infinite till break
        fmt.Println("loop")
        break
    }

    for n := range 6 { 
        if n%2 == 0 {
            continue //skip to next iteration
        }
        fmt.Println(n)//prints 1,3,5
    }
}
