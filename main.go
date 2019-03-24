package main

import "fmt"

func main() {
	var s int
	var f string
	for c := true; c; c = (s != 5) { //do while loop for golang. initial condition always true
		getSelection(&s) //call getSelection function to scan standard input passing pointer s to function as argument
		switch s {       //switch statement for s
		case 1: //if s is 1
			f = "Apple"
		case 2: //if s is 2
			f = "Kiwi"
		case 3: //if s is 3
			f = "Melon"
		case 4: //if s is 4
			f = "Banana"
		case 5: //if s is 5
			fmt.Println("Terminating program")
		default:
			fmt.Println("Selected option is not available")
		}
		if s < 5 {
			fmt.Println("Your favorite fruit is", f)
		}
	}
}

func getSelection(s *int) {
	fmt.Println("Please select one of your favorite fruite\n 1.Apple\n 2.Kiwi\n 3.Melon\n 4.Banana\n 5.Terminate program")
	fmt.Scan(s)
}
