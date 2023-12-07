package main

// 批量定义多个全局变量，小写开头或者"_"开头代表变量只能在当前包内被引用
var (
	globalNum1 int
	globalNum2 int = 0
	globalNum3     = 0
)

func main() {
	variables()
	variablesMulti()
}

func variables() {
	// 变量定义
	var num1 int
	var num2 int = 0
	var num3 = 0
	num4 := 0

	// 局部变量未使用会出现编译错误，如：num1 declared and not used
	// 忽略变量未使用的编译错误
	_ = num1
	_ = num2
	_ = num3
	_ = num4

	// 变量赋值
	num1 = 1
	num2 = 2
	num3 = 3
	num4 = 4

	// 变量使用
	println(num1)
	println(num2)
	println(num3)
	println(num4)
}

func variablesMulti() {
	// 在一个分组中批量定义多个变量
	var (
		num1 int
		num2 int = 0
		num3     = 0
	)

	// 使用简化形式批量定义多个变量并赋予初始值
	num4, num5 := 4, 5

	// 忽略编译错误
	_ = num1
	_ = num2
	_ = num3
	_ = num4
	_ = num5
}
