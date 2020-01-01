package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", HandleRoot)
	http.ListenAndServe(":8080", nil)
}

// HandleRoot : handle API request at /
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path[1:]
	number, err := strconv.Atoi(path)
	fmt.Println("number is ", number)
	if err != nil {
		fmt.Fprintf(w, "Please enter a number\n")
		return
	}
	var isTwoSidedPrime = IsDoubleSidedPrime(path)
	if isTwoSidedPrime {
		fmt.Fprintf(w, "Yes %d is a two sided prime\n", number)
	} else {
		fmt.Fprintf(w, "Sorry, %d is not a two sided prime :(\n", number)
	}
}

// IsPrime : checks whether given number is prime or not
func IsPrime(num int) bool {
	out := true
	sqrt := int(math.Ceil(math.Sqrt(float64(num))))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 && num != i {
			fmt.Println(num, " is divisible by ", i)
			out = false
			break
		}
	}
	var outString = ""
	if out {
		outString = "Prime"
	} else {
		outString = "Not Prime"
	}
	fmt.Println("number ", num, " is: ", outString, ", sqrt: ", sqrt)
	return out
}

// IsDoubleSidedPrime : main func to check two sided prime behaviour
func IsDoubleSidedPrime(numStr string) bool {
	totalSize := len(numStr)
	//checking left sided primality
	fmt.Println("checking left sided primality")
	for i := 0; i < totalSize; i++ {
		currentNumStr := numStr[i:]
		currentNum, err := strconv.Atoi(currentNumStr)
		if err != nil {
			fmt.Println("Unable to convert string to integer:", currentNumStr)
			return false
		}
		if !IsPrime(currentNum) {
			return false
		}
		fmt.Println(currentNum)
	}

	//checking right sided primality
	fmt.Println("checking right sided primality")
	for i := totalSize - 1; i > 0; i-- {
		currentNumStr := numStr[:i]
		currentNum, err := strconv.Atoi(currentNumStr)
		if err != nil {
			fmt.Println("Unable to convert string to integer:", currentNumStr)
			return false
		}
		if !IsPrime(currentNum) {
			return false
		}
		fmt.Println(currentNum)
	}
	return true
}
