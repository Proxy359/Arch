package main
import "fmt"

func first (n int) string {
y:="*";
sum:= "";
for i:=1;i<=n; i++ {
sum+=y
}
return sum
}

func main () {
fmt.Println(first(4))
}