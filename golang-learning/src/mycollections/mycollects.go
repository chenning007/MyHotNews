package mycollections

import "fmt"

func init() {

}

func MyCollects() {

	// 1. 数组专声明和初始化
	var array [5]int
	var array0 = [5]int{}
	var array1 = [5]int{0, 1, 2, 3, 4}
	var array2 = [...]int{0, 1, 2, 3, 4}
	var array3 = [5]int{1: 1, 2: 2}
	fmt.Println(array)               //[0 0 0 0 0]
	fmt.Println(array0, len(array0)) // [0 0 0 0 0] 5
	fmt.Println(array1, len(array1)) // [0 1 2 3 4] 5
	fmt.Println(array2, len(array2)) // [0 1 2 3 4] 5
	fmt.Println(array3, len(array3)) // [0 10 20 0 0] 5

	// 2. 使用数组
	// 2.1 访问数组元素
	fmt.Println(array2[2]) // 2
	array2[2] = -2
	fmt.Println(array2[2]) // -2

	var array4 = [5]*int{0: new(int), 1: new(int)}
	fmt.Println(array4) // [0xc00006e018 0xc00006e020 <nil> <nil> <nil>]
	*array4[0] = 10
	// *array4[2] = 20 // nil exception
	*array4[0] = 20
	fmt.Println(*array4[0], array4) // 20 [0xc000018070 0xc000018078 <nil> <nil> <nil>]

	// 数组复制
	var array5 [2]string
	var array6 = [2]string{"red", "blue"}
	fmt.Println(array6) // [red blue]
	fmt.Println(array5) // [ ]
	array5 = array6
	fmt.Println(array5) // [red blue]
	array5[1] = "yellow"
	fmt.Println(array5) // [red yellow]
	fmt.Println(array6) // [red blue]

}
