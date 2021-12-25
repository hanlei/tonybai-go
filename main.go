package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("17｜复合数据类型：用结构体建立对真实世界的抽象")

	fmt.Println("### 定义类型")

	// 底层类型(Underlying Type)，基于 Go 原生类型定义
	// 底层类型，是用来判断两个类型是否相同！
	// 底层类型相同，变量可以通过显示转换相互赋值
	type T1 int
	type T2 T1
	type T3 string

	var n1 T1
	var n2 T2 = 5
	//n1 = n2 不能直接赋值，因为类型不一样
	n1 = T1(n2)

	//var s T3 = "hello"
	//n1 = T1(s) // cannot convert
	fmt.Println(n1, n2)

	// 使用别名的方式
	type T = string

	var s string = "hello"
	var t T = s
	fmt.Println(t, "使用别名，其实就是一个类型，赋值不需要进行显示的类型转换")

	fmt.Println("### 定义结构体")

	type Book struct {
		Title   string         //书名
		Pages   int            //页数
		Indexes map[string]int //索引
		hide    string         //隐藏字段，不暴露给其他包
	}
	var b Book
	b.Title = "go语言从入门到放弃"
	b.Pages = 800

	fmt.Println(b)

	fmt.Println("### 定义一个空结构体")

	type Empty struct{}
	var em Empty
	{
	}
	fmt.Println(unsafe.Sizeof(em))

	fmt.Println("### 使用其他结构定义结构体中字段类型")

	type Person struct {
		Name  string
		Age   int
		Phone string
	}

	type Booken struct {
		Title  string
		Author Person
	}

	var booken Booken
	fmt.Println(booken.Author.Name)

	fmt.Println("以嵌入字段(Embedded Field),匿名字段方式定义结构")

	type Booken1 struct {
		Title string
		Person
	}

	var booken1 = Booken{
		Title: "1111",
		Author: Person{
			Name:  "hanlei",
			Age:   18,
			Phone: "139",
		},
	}
	var booken2 = Booken1{
		Title: "1111",
		Person: Person{
			Name:  "hanlei",
			Age:   18,
			Phone: "139",
		},
	}

	fmt.Println(booken1, booken2)
	fmt.Println(booken2.Name)
	fmt.Println(booken2.Person.Name)

	type Ttest struct {
		b byte
		i int64
		u uint16
	}

	type Stest struct {
		b byte
		u uint16
		i int64
	}

	var ttest Ttest
	println(unsafe.Sizeof(ttest)) // 24
	var stest Stest
	println(unsafe.Sizeof(stest)) // 16

	fmt.Println("基于内存对齐原则，结构体字段顺序不同，对内存空间的占用是不同的")

}
