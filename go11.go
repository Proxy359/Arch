package main

import "fmt"

func first (n int) string {
    y:="*";z:="@"; sum:="";
    for x:=1;x<=n;x++{
    for i:=1;i<=x;i++{
        if x%2==0 {sum+=y
	}else  {sum+=z}}
        sum+="\n"
    }
    return sum
}

func main () {
    fmt.Println(first(10))
}