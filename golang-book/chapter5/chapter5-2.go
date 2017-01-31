package main
 
import "fmt"
 
func main() {
    i := 1
    for i <= 10 {
		if i % 2 == 0 {
			// 짝수
			fmt.Println(i, "짝수")
		} else {
			// 홀수
			fmt.Println(i, "홀수")
		}

		switch i {
		case 0: fmt.Println("영")
		case 1: fmt.Println("일")
		case 2: fmt.Println("이")
		case 3: fmt.Println("삼")
		case 4: fmt.Println("사")
		case 5: fmt.Println("오")
		default: fmt.Println("알 수 없는 숫자")
		}
        
        i = i + 1
    }
}