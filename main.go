package main
import (
    "fmt" 
    "net/http"
    "strconv"
)


func main(){
    http.HandleFunc("/",IsItADoubleSidedPrime)
    http.ListenAndServe(":8080",nil)
}


func IsItADoubleSidedPrime(w http.ResponseWriter, r *http.Request){
    //https://en.wikipedia.org/wiki/Truncatable_prime#Lists_of_truncatable_primes
    TwoSidedPrimes:= [15]int{2, 3, 5, 7, 23, 37, 53, 73, 313, 317, 373, 797, 3137, 3797, 739397}
    var path = r.URL.Path[1:]
    number,err := strconv.Atoi(path)
    fmt.Println("number is ",number)
    if err!=nil {
        fmt.Fprintf(w,"Please enter a number\n")
        return
    }
    for i:=0;i<15;i++{
        if(number == TwoSidedPrimes[i]){
            fmt.Fprintf(w, "Yes %d is a two sided prime\n",number)
            return
        }
    }
    fmt.Fprintf(w,"Sorry, %d is not a two sided prime :(\n", number)
}
