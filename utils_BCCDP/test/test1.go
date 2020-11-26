package main

func main() {
	//myNumber := MyNumber{
	//	name: "王二狗",
	//	num:  20,
	//}
	//yourNumber := YourNumber{
	//	name: "李小花",
	//	num:  30,
	//}
	//yourNumber.sub(30,20)
}

type IsNumber interface {
	add(int, int) int
	sub(int, int) int
}
type IsFloat interface {
	add(int, int) int
	sub(int, int) int
}
type MyNumber struct {
	name string
	num  int
}
type YourNumber struct {
	name string
	num  int
}

func (*YourNumber) add(a int, b int) int {
	return a + b
}
func (*YourNumber) sub(a int, b int) int {
	return a - b
}
func (*MyNumber) add(a int, b int) int {
	return a * b
}
func (*MyNumber) sub(a int, b int) int {
	return a / b
}