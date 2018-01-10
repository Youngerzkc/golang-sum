package main

import (
	"fmt"
)

func main() {
	fmt.Println("a return:", a()) // 打印结果为 a return: 0
}

func a() int {
	var i int
	defer func(){
		fmt.Println("i",i)
	}()
	defer func() {
		i++
		fmt.Println("a defer2:", i,&i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i,&i) // 打印结果为 a defer1: 1
	}()
	fmt.Println("i====",i,&i)
	return i
}
//a()int 函数的返回值没有被提前声名，其值来自于其他变量的赋值，而defer中修改的也是其他变量，
//而非返回值本身，因此函数退出时返回值并没有被改变。