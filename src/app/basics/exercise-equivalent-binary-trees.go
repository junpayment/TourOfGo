package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if nil != t.Left {
			walker(t.Left)
		}

		ch <- t.Value

		if nil != t.Right {
			walker(t.Right)
		}
	}
	walker(t)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, res1 := <-ch1
		v2, res2 := <-ch2

		if !(v1 == v2 && res1 == res2) {
			return false
		}

		if !res1 {
			break
		}
	}
	return true
}

func main() {
	//ch := make(chan int)
	//go Walk(tree.New(1), ch)
	//
	//var i int
	//var ok bool
	//
	//for {
	//	i, ok = <-ch
	//	if ok {
	//		fmt.Println(i)
	//	}
	//	time.Sleep(time.Second)
	//}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(5), tree.New(5)))
	fmt.Println(Same(tree.New(3), tree.New(2)))
}
