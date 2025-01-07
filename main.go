package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

var errRequestFailed = errors.New("Request failed")
var results = make(map[string]string)

var urls = []string{
	"https://www.google.com",
	"https://www.naver.com",
	"https://www.daum.net",
	"https://www.github.com",
	"https://www.youtube.com",
	"https://www.facebook.com",
	"https://www.instagram.com",
	"https://www.twitter.com",
}

func main() {
	const name string = "Gopher"

	//fmt.Println(multiply(2, 3))
	//
	//fmt.Println(lenAndUpper(name))
	//
	//totalLength, upperName := lenAndUpper(name)
	//
	//fmt.Println(totalLength, upperName)
	//
	//repeatMe("nico", "lynn", "dal", "marl", "flynn")

	//fmt.Println(nakedReturn("zaur"))
	//
	//deferTest()
	//
	//fmt.Println(superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	//
	//fmt.Println(canIDrink(16))
	//
	//fmt.Println(canIDrinkWithSwitch(16))

	//pointerTest()

	//arrayTest()

	//mapTest()

	//structTest()

	//account := accounts.NewAccount("zaur")
	//account.Deposit(10000)
	//err := account.Withdraw(1500)
	//if err != nil {
	//	//log.Fatal은 프로그램을 종료시킴
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(account)
	//fmt.Println(account.Balance(), account.Owner())

	//mydict := dict.Dictionary{}
	//mydict["first"] = "First"
	//
	//definition, err := mydict.Search("first2")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(definition)
	//}
	//
	//err = mydict.Add("first", "Greeting")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//mydict.Update("first", "World")
	//
	//mydict.Delete("first")
	//fmt.Println(mydict)

	//for _, url := range urls {
	//	result := "OK"
	//	err := hitURL(url)
	//	if err != nil {
	//		result = "FAILED"
	//	}
	//	results[url] = result
	//
	//}
	//
	//fmt.Println(results)
	c := make(chan string)
	people := []string{"zaur", "donghee", "flynn", "marl", "dal"}
	for _, person := range people {
		go isSexy(person, c)
	}

	//fmt.Println("waiting for messages")
	//// 채널은 메시지를 받을 때까지 기다림 그냥 먼저 오는거 받음
	//fmt.Println("receive this message", <-c)
	//fmt.Println("receive this message", <-c)

	for _, val := range people {
		fmt.Print("waiting for ", val, " ")
		fmt.Println(<-c)
	}
}

func multiply(a int, b int) int {
	return a * b
}

// 리턴 값이 여러개일 때
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 리턴 값의 이름을 정해줄 수 있음 이럴 땐 naked return 이 가능함
func nakedReturn(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// defer는 함수가 끝난 후 실행됨
func deferTest() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")
}

// ...string은 여러개의 string을 받을 수 있음
func repeatMe(words ...string) {
	fmt.Println(words)
}

// for range
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// if 문 안에서 변수를 만들 수 있음 그리고 ;로 구분
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

// java와 달리 case문에 조건을 넣을 수 있음
func canIDrinkWithSwitch(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	}
	return false
}

// &는 메모리 주소를 보여줌 *는 메모리 주소에 있는 값을 보여줌
func pointerTest() {
	a := 2
	b := &a
	a = 10
	fmt.Println(*b)
}

func arrayTest() {
	names := []string{"nico", "lynn", "dal"}
	name1 := &names
	names = append(names, "flynn")
	names2 := &names
	fmt.Println(name1 == names2)

	fmt.Println(names)
}

// []안의 타입은 key의 타입이고 []뒤의 타입은 value의 타입
func mapTest() {
	nico := map[string]string{"name": "nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	}

	fmt.Println(nico["name"])
}

func structTest() {
	type person struct {
		name    string
		age     int
		favFood []string
	}

	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}

// main() 은 goroutine을 기다리지 않고 바로 종료됨
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

// 채널은 채널 타입 뒤에 메인 함수로 보낼 타입도 적음
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
