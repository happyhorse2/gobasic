package basictype

import (
	"fmt"
)

// Slicelencap :切片数据结构为，指向存储结构的指针，从指针寻找的长度，存储结构的容量
func Slicelencap() {
	a := []int{1, 2, 3, 4} //初始化切片
	fmt.Println("a:", len(a), cap(a), a)

	b := [10]int{1, 2, 3, 4} //初始化数组
	fmt.Println("b:", len(b), cap(b), b)

	c := make([]int, 4, 10) //初始化切片
	fmt.Println("c:", len(c), cap(c), c)

	d := b[:10] //初始化切片
	fmt.Println("d:", len(d), cap(d), d, &d, &b)

	e := append(d, 5) //append后d的容量不变
	fmt.Println("e:", &e)
	e[0] = 100 //没超出底层数组的容量，因此e和d都指向同一个数组，修改e会影响d,  超出数组b的容量后，e会指向新的数组
	fmt.Println("d after append:", len(d), cap(d), d)
	fmt.Println("e:", len(e), cap(e), e, &d, &e)
	fmt.Println("b:", len(b), cap(b), b) // 查看数组
}

// Slicelencapnil :空切片
func Slicelencapnil() {
	// 切片默认初始化为nil, 长度为0，长度为0，指向存储结构的指针是空，存储容量，切片长度均初始化为0
	var a []int
	fmt.Print("a:", len(a), cap(a), a)
	b := make([]int, 0)
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}
}

// Slicelencaptwo :切片扩容练习
func Slicelencaptwo() {
	path := []int{1, 2, 3, 4, 5}
	t1 := path[2:3]
	fmt.Println(len(t1), cap(t1), len(path), cap(path), t1[0])
	t1 = append(t1, []int{10, 11, 13}...)
	fmt.Println(len(t1), cap(t1), len(path), cap(path))
	fmt.Println(path)
	fmt.Printf("%p %p %p %p %p\n", path, t1, &path[0], &t1[0], &path[2])
	fmt.Printf("%p %p\n", &path, &t1)
	t1 = append(t1, []int{10, 12, 13}...)
	fmt.Println(path)
	fmt.Printf("%p %p\n", path, t1)
	fmt.Printf("%p %p\n", &path, &t1)
}
