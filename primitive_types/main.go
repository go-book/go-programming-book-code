package main

func main() {
	ok := true
	println(ok)

	var uint8N uint8 = 255
	println(uint8N)

	var uint16N uint16 = 65535
	println(uint16N)

	var uint32N uint32 = 4294967295
	println(uint32N)

	var uint64N uint64 = 18446744073709551615
	println(uint64N)

	var int8N int8 = 127
	println(int8N)

	var int16N int16 = 32767
	println(int16N)

	var int32N int32 = 2147483647
	println(int32N)

	var int64N int64 = 9223372036854775807
	println(int64N)

	var intN int = 1
	println(intN)

	var uintN uint = 1
	println(uintN)

	var uintptrN uintptr = 1
	println(uintptrN)

	var byteN byte = 255
	println(byteN)

	var runeN rune = '中'
	println(runeN)

	var f32 float32 = 1.1
	println(f32)

	var f64 float64 = 1.1
	println(f64)

	// 创建complex64类型的复数
	c64 := complex(float32(1.1), float32(1.2))
	// 获取实部
	real64 := real(c64)
	// 获取虚部
	imag64 := imag(c64)
	println(c64)
	println(real64)
	println(imag64)

	// 创建complex128类型的复数
	c128 := complex(1.1, 1.2)
	// 获取实部
	real128 := real(c128)
	// 获取虚部
	imag128 := imag(c128)
	println(c128)
	println(real128)
	println(imag128)

	s := "hello world"
	println(s)
	// 将字符串转为byte切片
	bytes := []byte(s)
	println(bytes)
	// 将byte切片转为字符串
	s2 := string(bytes)
	println(s2)
	// 定义一个字符
	var char rune = '文'
	// 将字符转为字符串
	s = string(char)
	println(s)
}
