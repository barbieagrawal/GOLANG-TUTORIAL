package main
import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println("Hello World!")

	var x int32 = 14 //defining variables with initialization
	fmt.Println(x)

	var y float32 = 10.2
	fmt.Println(y)

	var prod float32 = float32(x) * y //typecast for arithmatic operations
	fmt.Println(prod) 

	var s string = "Hello" + " " + "World"
	fmt.Println(s)

	fmt.Println(len("#")) //len -> length in bytes
	fmt.Println(utf8.RuneCountInString("#")) //length in characters

	var myRune rune = 'a'
	fmt.Println(myRune) //prints ascii

	var i int32 //initialization not compulsory
	fmt.Println(i) //prints default value (for int = 0)
	
	//value for constants can't be changed and have to be initialized
	const c string = "const val"
	fmt.Println(c)

	printMe() //function call 

	var num int = 11 //intDivision function call
	var den int = 2
	var result, remainder int = intDivision(num, den)
	fmt.Printf("The result of integer division is %v with remainder %v\n", result, remainder)

	switch remainder { //conditional switch statements
	case 0:
		fmt .Printf("The division was exact")
	case 1,2: 
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
}

func printMe() {
	fmt.Println("ME!")
}

func intDivision(num int, den int) (int, int) {
	var result int = num / den
	var remainder int = num % den
	return result, remainder
}