package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
			println("hello,world")
			fmt.Println("hello,world")
			n := 100 + 200
			m := n + 100
			msg := "hoge" + "fuga"

			//if
			if n == 300 {
				fmt.Println(n)
			} else if n == 100 {
				fmt.Print(m)
			} else {
				fmt.Println(msg)
			}

			//switch
			switch n {
			case 100:
				fmt.Println("n=100")
			default:
				fmt.Println("default")
			}
			for i := 0; i < 100; i++ {
				print(i)
			}

			var x int
		LOOP:
			for {
				fmt.Println(x)
				x++
				if x == 100{
					break LOOP
				}

			}
	*/
	// 100をHex型として代入
	var hex Hex = 100
	// Stringメソッドを呼び出す
	fmt.Println(hex.String())

}
func main2() {
	for x := 0; x < 100; x++ {
		if x%2 == 0 {
			fmt.Println(x, ":even")
		} else {
			fmt.Println(x, ":odd")
		}
	}
}
func main3() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	n := rand.Intn(6)

	switch n {
	case 6:
		fmt.Println("大吉")
	case 5, 4:
		fmt.Println("中吉")
	case 3, 2:
		fmt.Println("小吉")
	default:
		fmt.Println("凶")
	}
}

func main4() {
	var sum int
	sum = 5 + 6 + 3
	var avg float32 = (float32(sum / 3))
	if avg > 4.5 {
		println("good")
	}
}

func main5() {
	var a, b, c bool
	if a && b || !c {
		println("True")
	} else {
		println("false")
	}
}
func main6() {
	p := struct {
		name string
		age  int
	}{
		name: "Gopher",
		age:  10,
	}
	p.age++
	fmt.Println(p.name, p.age)

	var ns = [5]int{1, 2, 3, 4, 5}
	println(ns[3])
	println(len(ns))
	println(ns[1:3])
	var slice = []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	m := map[string]int{"x": 10, "y": 20}
	//println(m["x"])
	m["z"] = 30
	//println(m["z"])

	if n, ok := m["s"]; ok {
		fmt.Println("n:", n)
		fmt.Println("ok:", ok)
	} else {
		fmt.Println("n:", n)
		fmt.Println("ok:", ok)
	}
	type Person struct {
		Name string
	}
	var test_name Person
	test_name.Name = "aaa"
	test_name2 := Person{Name: "bbb"}
	fmt.Println(test_name2.Name)
}

func add(x int, y int) int {
	return x + y
}
func swap(x, y int) (int, int) {
	return y, x
}
func swap2(x, y int) (x2 int, y2 int) {
	y2, x2 = x, y
	return y2, x2
}

func main7() {
	p := struct {
		age  int
		name string
	}{age: 10, name: "Gopher"}
	p2 := p // コピー
	p2.age = 20
	println(p.age, p.name)

	var x int = 100
	var xp *int
	xp = &x
	print(*xp)

}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}
