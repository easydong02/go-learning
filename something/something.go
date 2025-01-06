package something

import "fmt"

// 소문자로 시작하면 다른 패키지에서 사용 불가능
func sayBye() {
	fmt.Println("Goodbye!")
}

// 대문자로 시작하면 다른 패키지에서 사용 가능
func SayHello() {
	fmt.Println("Hello!")
}
