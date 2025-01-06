package main
import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println("Hello World!")

	var x int32 = 14 //defining variables with initialization
	fmt.Println(x)

	var y float32 = 10.2
	fmt.Println(y)

	var prod float32 = float32(x) * y //typecast for arithmetic operations
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

	//conditional switch statements

	switch remainder { 
	case 0:
		fmt .Printf("The division was exact\n")
	case 1,2: 
		fmt.Printf("The division was close\n")
	default:
		fmt.Printf("The division was not close\n")
	}

	//arrays 

	var arr [3]int32 //by default array elements initialized to 0
	arr[1] = 123
	fmt.Println(arr[0]) //first element 
	fmt.Println(arr[1:3]) //second and third elements 

	fmt.Println(&arr[0]) //location in continguous memory alloc of array
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	var arr2 [3]int32 = [3]int32{1, 2, 3} //a way of initialization
	fmt.Println(arr2[0:3])

	arr3 := [...]int32{4, 5, 6}
	fmt.Println(arr3)

	//slices
	//arrays have fixed size, slices can have variable sizes

	var sl []int32 = []int32{7, 8, 9}
	fmt.Println(sl)
	fmt.Printf("The length is %v with capacity %v\n", len(sl), cap(sl))
	sl = append(sl, 10)
	fmt.Println(sl)
	fmt.Printf("The length is %v with capacity %v\n", len(sl), cap(sl))

	//maps

	var m map[string]uint8 = make(map[string]uint8)
	fmt.Println(m)

	var m2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(m2["Adam"])
	fmt.Println(m2["Barb"]) //by deafult 0 value is printed

	//for loop 

	for name, age := range m2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for i, v := range arr2 {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	//strings, runes, bytes

	var s1 = "résumé"
	var indexed = s1[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range s1 {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of s1 is %v\n", len(s1))
}

func printMe() {
	fmt.Println("ME!")
}

func intDivision(num int, den int) (int, int) {
	var result int = num / den
	var remainder int = num % den
	return result, remainder
}