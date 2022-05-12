package main

import "fmt"

func main () {
    fmt.Println (fan(4))
}

func fan (n int) int{
    return n*n
}
