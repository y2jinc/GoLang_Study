package main
 
import "fmt"
 
const x string = "Hello World"
var (
    a = 5
    b = 10
    c = 15
)
 
func f() {
	x := 3
    fmt.Println(x)
}

func input_test() {
	fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

	output := input * 2
    fmt.Println(output)
}

func main() {
	fmt.Println("-----------------------------")
    fmt.Println(x)
	
	fmt.Println("-----------------------------")
	f()

	fmt.Println("-----------------------------")
	fmt.Println("a :", a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("-----------------------------")
	test_x := 5
	test_x += 1
	fmt.Println(test_x)

	fmt.Println("-----------------------------")
	input_test()
}