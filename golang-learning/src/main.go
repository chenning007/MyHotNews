package main

import (
	"fmt"
	"mycollections"
	"mypack" // 必需导入才能被使用
)

func main() {
	fmt.Println("start from golang learning...")

	fmt.Println("--------------- start --------------------") // 分割线

	fmt.Println("========== 1. 包相关实例 ==========")
	mypack.Myprint("hello golang")

	fmt.Println()

	fmt.Println("========== 2. 集合相关实例 ==========")
	// mycollects.ArraysExamples() // 2.1 数组
	// mycollects.SlicesExamples() // 2.2 切片
	mycollections.MappingExamples() // 2.3 映射

	fmt.Println()
	fmt.Println("--------------- end --------------------") // 分割线

	fmt.Println("end of golang learning...")
}
