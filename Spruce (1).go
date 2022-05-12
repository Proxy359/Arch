package main

import "fmt"

func first(n int) string {
y:="*";sum:="";
    for x:=1;x<=n;x++{
        for i:=1; i<=x; i++{
            sum+=y
        }
        sum+="\n"
    }
    return sum
}

func main () {
    fmt.Println (first(5))
}
