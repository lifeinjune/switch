package main

import "fmt"

func main() {
	var s int
	var f string
	for c := true; c; c = (s != 5) {
		getSelection(&s)
		switch s {
		case 1:
			f = "Apple"
		case 2:
			f = "Kiwi"
		case 3:
			f = "Melon"
		case 4:
			f = "Banana"
		case 5:
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
