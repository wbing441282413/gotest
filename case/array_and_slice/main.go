package main

import (
	"fmt"
	"sort"
)

func main() {

	//初识数组
	fmt.Println("------------------初始数组---------------")
	var aa [3]int = [3]int{1, 2, 3}
	cc := [...]int{1, 2} //只有数组有[...]的方式
	var dd [2]int        //默认初始是0
	fmt.Printf("%T\n", cc)
	fmt.Printf("%T\n", aa)
	for _, a := range dd {
		fmt.Println(a)
	}

	//二维数组
	fmt.Println("------------------二维数组---------------")
	k := [2][2]int{{1, 1}, {2, 2}}
	var j [2]int
	j = k[1] //将二位数组的一个维度赋值给一个一维数组
	for _, a := range j {
		fmt.Println(a)
	}

	//数组的默认值
	fmt.Println("------------------数组的默认值---------------")
	var gg [2]int
	for _, a := range gg {
		fmt.Println(a)
	}

	//初识切片
	fmt.Println("------------------初识切片---------------")
	slice := make([]int, 2, 3)
	for _, a := range slice {
		fmt.Println(a)
	}
	fmt.Printf("%p\n", slice) //可以看到切片是一个引用是个指针

	//切片和底层数组是够用的
	fmt.Println("------------------切片和底层数组共用内存---------------")
	ii := [3]int{1, 2, 3}
	aaa := ii[:] //运行结果看到切片并不会开辟新的内存空间
	fmt.Printf("%p \n", &ii)
	fmt.Printf("%p \n", aaa) //也可以看到切片是个引用，是个指针

	//range在切片中的使用
	fmt.Println("------------------range在切片中的使用---------------")
	ll := []int{1, 2, 3, 4, 5}
	l := ll[2:5]
	fmt.Println("从下标2开始，取5-2=3个元素")
	for _, a := range l {
		fmt.Println(a)
		a = 9 //range返回的a是切片元素的复制，而不是元素的引用，所以这里我们修改v的值并不会改变slice切片里的值
	}
	fmt.Println(l)

	//切片长度，容量
	fmt.Println("------------------切片长度，容量---------------")
	sl := []int{1, 2, 3, 4, 5}

	hscl := sl[1:2] //切片长度是j-i=1，只能访问切片的长度大小的元素，超过就要切片扩容
	// fmt.Println(hscl[1]) //比如这里就越界了
	fmt.Println(cap(hscl)) //可以看到没有指定k值的时候，k就是底层数组的个数,也就是5，于是切片容量cap是k-i=5-1=4

	newSlice := sl[1:2:4] //i,j,k,第三个参数是k,容量是k-i,就是4-1=3，切片长度是j-i = 2-1 =1
	fmt.Println(newSlice)
	// fmt.Println(newSlice[1]) //只能访问切片的长度大小的元素，会越界

	//切片append
	fmt.Println("------------------切片append---------------")
	fmt.Println("底层数组不够的时候，这时切片的容量肯定也不够用了，会新建一个新的数组：")
	sll := append(newSlice, 10, 11, 12, 13, 14, 15)
	fmt.Println("sll----", sll)
	fmt.Println("newSlice----", newSlice)
	fmt.Println("sl----", sl) //可以看到原底层数组和newSlice都没变
	fmt.Printf("%p\n", sll)
	fmt.Printf("%p\n", newSlice) //可以看到地址不一样了，是新建了一个底层数组的
	//这里比较newSlice地址和sxl就行了，和比较底层数组是一样的

	fmt.Println("当切片的容量还够，追加后切片长度不超过切片的容量，不会新建数组：")
	ssl := append(newSlice, 10)
	fmt.Println("ssl----", ssl)
	fmt.Println("newSlice----", newSlice)
	fmt.Println("sl----", sl) //可以看到原底层数组元素值变了
	fmt.Printf("%p\n", ssl)
	fmt.Printf("%p\n", newSlice) //可以看到地址还是一样，没有新建数组
	//这里比较newSlice地址和sxl就行了，和比较底层数组是一样的

	/**
	因为切片的容量不够的时候，会直接新建一个新的底层数组然后把原来数组中的值都复制到新数组中，
	所以原数组的值是不会变的
	注意：判断的标准是切片的容量，而不是底层数组是否还有空间，只要是追加后切片长度超过切片的容量，
	不管底层数组是否有空间都会去新建一个新的数组复制元素
	*/
	fmt.Println("底层数组够,但是追加后超过切片的容量了，同样新建一个底层数组的：")
	sxl := append(newSlice, 10, 11, 12)
	fmt.Println("sxl----", sxl)
	fmt.Println("newSlice----", newSlice)
	fmt.Println("sl----", sl)
	fmt.Printf("%p\n", sxl)
	fmt.Printf("%p\n", newSlice) //可以看到地址不一样了，是新建了一个底层数组的
	//这里比较newSlice地址和sxl就行了，和比较底层数组是一样的

	newSlice = append(newSlice, sl...)
	fmt.Println(newSlice)
	fmt.Println(sl)

	for _, a := range sl {
		a = a + 1
	}
	fmt.Println(sl) //可以看到没有变

	//排序切片
	sort.Ints(newSlice)
	fmt.Println(newSlice)
}
