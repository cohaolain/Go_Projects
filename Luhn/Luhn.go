package main

import "fmt"
import "strconv"

func main() {
	
	fmt.Println("\n    Welcome to Ciarán Ó hAoláin's Luhn Algorithm verification program, written in Go!\n\n    Start by typing a number, and hitting return!\n")	
	
	for {	
		var number int
		fmt.Scanf("%d", &number)
		
		if number!=0{
		
			if isValid(number) {
				fmt.Println("Valid\n\n")
			} else {
				fmt.Println("Invalid\n\n")
			}
			
		}		
	}
	
}

func isValid(num int) bool {

	length := (len(strconv.Itoa(num)))

	// Create array containing each digit of the number
	array := make([]int, length)
	for i:=0;i<length;i++ {
		array[i]=num%10
		num/=10
	}

	/* Print the number as it's stored in the array
	for i:=length-1;i>=0;i-- {
		fmt.Print(array[i])
	}
	fmt.Println()
	*/	
	
	// Apply the algorithm
	count := 0
	for i:=0;i<len(array);i+=2{
		count+=array[i]
	}
	for i:=1;i<len(array);i+=2{
		array[i]*=2
		if array[i]>9 {
			array[i]-=9
		}
		count+=array[i]
	}
	
	// Run final check and print the results
	if count%10==0 {
		return true
	} else {
		return false
	}

}