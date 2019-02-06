package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"testing"
)

func TestHelloPointer(t *testing.T) {
	a := 5
	b := &a
	c := &a
	fmt.Printf("the address of a is %p \n", b)
	fmt.Printf("the address of a is %p \n", c)
	fmt.Printf("the value of a is %d \n", a)
	*c = 7
	fmt.Printf("the value of a is %d \n", a)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHelloWeb(t *testing.T) {

	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

type Shape interface {
	area() float64
}

type Abstract struct {
	shape Shape
}

func TestHelloMuti(t *testing.T) {
	a := &Abstract{
		shape: &Circle{x: 0, y: 0, radius: 100},
	}

	fmt.Println(a.shape.area())
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area();
}
func TestHelloInterface(t *testing.T) {
	c := &Circle{x: 0, y: 0, radius: 100}
	r := &Rectangle{width: 100, height: 10000}

	fmt.Printf("the circle area is %f and the rectangle area is %f \n", getArea(c), getArea(r))
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

type Person struct {
	name string
	age  int
}

func (p Person) sayHello() string {
	return "Hello my name is " + p.name + "and my age is " + strconv.Itoa(p.age)
}

func TestHelloStruct(t *testing.T) {
	p := Person{name: "药耀源", age: 26}
	fmt.Println(p)
	fmt.Println(p.sayHello())
}

func TestHelloClourse(t *testing.T) {
	sum := adder()

	for i := 0; i < 100; i++ {
		fmt.Println(sum(i))
	}
}

func TestHelloRange(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 5, 6, 6, 7, 7}

	for index, num := range nums {
		fmt.Printf("the index of %d is %d \n", index, num)
	}
}

func TestHelloMap(t *testing.T) {
	emails := make(map[int]int)
	for i := 1; i <= 100; i++ {
		emails[i] = i
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(emails[i])
	}

	fmt.Println(emails)
}

func TestHelloGo(t *testing.T) {
	println(HelloGo())

	a := make(chan bool, 1)
	b := make(chan bool)
	Exit := make(chan bool)

	//  <- chan 当chan中没有数据时会阻塞
	//  a <- true 往chan a 中放入放入一个值,此时阻塞的协程开始执行
	//  go func 启动一个协程
	go func() {
		for i := 1; i < 100; i += 2 {
			if ok := <-a; ok {
				fmt.Println(i)
				b <- true
			}
		}

	}()

	go func() {
		for i := 2; i <= 100; i += 2 {
			if ok := <-b; ok {
				fmt.Println(i)
				a <- true
			}
		}
		// 循环打印结束,主协程退出
		close(Exit)
	}()

	fmt.Println("end")

	a <- true
	<-Exit
}
