package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type Hero struct {
	Name string
	Age  int
}

type HelloSlice []Hero

func (hs HelloSlice) Len() int {
	return len(hs)
}

func (hs HelloSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HelloSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

type Monkey struct {
	Name string
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, " can climb tree!!!")
}

type LittleMonkey struct {
	Monkey
}

type BirdAble interface {
	Flying()
}

func (monkey *LittleMonkey) Flying() {
	fmt.Println(monkey.Name, "CAN FLY...")
}

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Usb
	Name string
}

func (phone Phone) Start() {
	fmt.Println("phone start ...")
}

func (phone Phone) Stop() {
	fmt.Println("phone start ...")
}

func (phone Phone) Call() {
	fmt.Println(phone.Name, "phone calling ...")
}

func main() {

	if result, errorMsg := divide(100, 10); errorMsg == "" {

		fmt.Println(result)
	}
	go say("world")
	say("hello")
	fmt.Print(num)

	var mapping = make(map[string]string)
	mapping["name"] = "liyuanhang"
	mapping["age"] = "28"
	for s, s2 := range mapping {
		fmt.Println(s, s2)
	}

	for s := range mapping {
		println(mapping[s])
	}

	var heroes HelloSlice

	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	for i, hero := range heroes {
		fmt.Println(i, hero)
	}

	sort.Sort(heroes)
	fmt.Println("排序后————————————————————")
	for _, v := range heroes {
		fmt.Println(v.Name, v.Age)
	}

	monkey := LittleMonkey{
		Monkey{
			Name: "WUKONG",
		},
	}
	monkey.climbing()
	monkey.Flying()

	var phone Phone = Phone{Name: "xiaomi"}
	//phone := Phone{Name: "xiaomi"}
	Working(phone)

	//闭包
	f2 := makeSuffix(".jpg")
	fmt.Println(f2("winter"))
	fmt.Println(f2("summer.jpg"))
}

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		if !(strings.HasSuffix(name, suffix)) {
			name = name + suffix
		}
		return name
	}
}

func Working(usb Usb) {
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func divide(dividee int, divider int) (result int, errorMessage string) {
	if divider == 0 {
		dataError := DivideError{
			dividee: dividee,
			divider: divider,
		}
		errorMessage = dataError.error()
		return
	} else {
		return dividee / divider, ""
	}
}
