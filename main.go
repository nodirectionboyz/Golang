package main
import "fmt"

func get_prime_factors(n int) []int {
	var slice []int
	var result int

	for n > 1 {
		result = small_factor(n)
		slice = append(slice, result)
		n = n/small_factor(n)
	}
return slice
}


func small_factor(n int) int {
	x := 2
	for i:=2; n % i != 0; i ++ {
		x = i + 1
	}
return x
}


func max(slice []int) int {
	var largest int
	for _, v := range slice {
		if v > largest {
			largest = v
		}
	}
	return largest
}


func main(){
	var n int
	fmt.Print("Please enter a number:")
	fmt.Scan(&n)
	slice := get_prime_factors(n)
	largest := max(slice)
	fmt.Println("Array of prime factors:",slice)
	fmt.Println("largest Prime Factor:",largest)
}

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
