//package main
//
//var a = "G"
//
//func main() {
//	n()
//	m()
//	n()
//}
//
//func n() { print(a) }
//
//func m() {
//	a := "O"
//	print(a)
//}

//package main
//
//var a = "G"
//
//func main() {
//	n()
//	m()
//	n()
//}
//
//func n() {
//	print(a)
//}
//
//func m() {
//	a = "O"
//	print(a)
//}

package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}