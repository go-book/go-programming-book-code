package main

// 批量定义多个全局常量，小写开头或者"_"开头代表变量只能在当前包内被引用
const (
	globalNum1 int = 0
	globalNum2     = 0
)

func main() {
	constants()
	constantsMulti()
}

func constants() {
	// 常量定义
	const num1 int = 0
	const num2 = 0

	// 常量不使用也不会出现编译错误

	// 常量不可二次赋值，否则会导致编译错误，例如：cannot assign to num1 (neither addressable nor a map index expression)
	//num1 = 1
	//num2 = 2

	// 常量使用
	println(num1)
	println(num2)
}

func constantsMulti() {
	// 在一个分组中批量定义多个常量
	const (
		num1 int = 0
		num2     = 0
	)
}
