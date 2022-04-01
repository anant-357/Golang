
package main
import "fmt"

func main(){
    var t,a int
    var x []int 
    fmt.Scan(&t)
    for t>0{
        fmt.Scan(&a)
        x = append(x,a)
        t--
    }
    fmt.Printf("%v",x)
    s:=0
    cnt:=0
    var i int = 0
    eatornot(x,cnt,s,i)
    a=len(x)
    fmt.Printf("%v",cnt)
}

func eatornot(x []int,cnt int,s int,i int){
    fmt.Printf("\n%v %v %v",cnt,s,i)
    if x[i]>0 {eatornot(x,cnt+1,s+x[i],i+1)}else{ if s+x[i]>0{for j:=0;j<2;j++{if j==0{eatornot(x,cnt,s,i+1)}else{eatornot(x,cnt+1,s+x[i],i+1)}}}}
    i++
    
    
}
