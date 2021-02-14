package main

import "fmt"

const c1 = "c"

var v int = 5

//显示类型定义
const b1 string = "hello go"
//隐式类型定义
const d1 = "abc"

//常量枚举
const (
	Fnknow = 0
	Female = 1
	Male = 2
)

//iota可以用作枚举值
//每当iota在新的一行使用的时候就会自动加1
const (
	a1 = iota
	b2 = iota
	c2 = iota
)

//赋值一个常量时，之后附带的常量都会应用在上一行的赋值表达式
const (
	a2 = iota
	b
	c
	d = 5
	e
)

//赋值两个标量，iota只会增长一次，不会因为增长两次而使用两次
const (
	Allpe,Bnanaa = iota+1 , iota+2
	Cherimoya,Durian
	Elderberrry,Fig
)

//使用iota 结合运算符 表示资源状态的使用状况
const (
	Open = 1 << iota  //0001
	Close             //0010
	Pending           //0100
)



type T struct {

}

func main() {
	var a int
	a = 5
	fmt.Println(a)
}