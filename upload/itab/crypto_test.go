package itab

import (
	"fmt"
	"testing"
)

func Test_generateHMAC(t *testing.T) {
	// 替换这里的值为你的密钥和消息
	// D5cZyknjFiESe01H0kCc36PJqiU=
	//res := convertArrayToString([]int{1, 2, 3, 4, 5, 6})
	//fmt.Println(res)
	// 示例用法
	//x := int32(1732584193)
	//y := int32(1655872086)
	//result := safeAdd(x, y)
	//fmt.Println(result)
	// 示例用法
	//value := uint32(12345)
	//shift := uint32(5)
	//result := leftRotate(value, shift)
	//fmt.Println(result)
	// 示例用法
	//tt := uint32(30)
	//b := uint32(123)
	//c := uint32(456)
	//d := uint32(789)
	//result := sha1F(tt, b, c, d)
	//fmt.Println(result)
	// 示例用法
	//tt := uint32(30)
	//result := sha1K(tt)
	//fmt.Println(result)
	// 示例用法
	//// 示例输入
	//message := make([]uint32, 3) // 举例使用长度为16的切片
	//length := 2                  // 举例使用长度为256的位数
	//message[0] = 1732584193 - 99
	//message[1] = 1732584193 - 88
	//message[2] = 1732584193 - 77
	//result := padInput(message, length)
	//fmt.Println(result)
	// 示例用法
	//inputStr := "Hello World!@#$"
	//result := uriEncodeDecode(inputStr)
	//fmt.Println("Hello%20World!%40%23%24")
	//fmt.Println(result)
	// 示例用法
	//inputStr := "abcd"
	//result := utf8ToByteArray(inputStr)
	//fmt.Println(result)
	fmt.Println(encryptOrHash("1", "1"))

	// 测试
	//input := "Hello, World!"
	//encoded := base64Encode(input, byte('='))
	//fmt.Println(encoded)
	// const result = hx.b64_hmac("5Y5b2hWisRnXXNUJOcPtkg1v2R9dZK", "eyJleHBpcmF0aW9uIjoiMjAyMy0xMi0wNVQwODowNzoyMS4xMjNaIiwiY29uZGl0aW9ucyI6W3siYnVja2V0IjoieGRsdW1pYTIifSx7ImtleSI6InVzZXItd2Vic2l0ZS1pY29uLzIwMjMxMjA0L1dOUTZPY19YcGNOdG11VGZxTk5INzA0ODUucG5nIn0sWyJjb250ZW50LWxlbmd0aC1yYW5nZSIsMCwxMDczNzQxODI0XV19");
	accessKeySecret := "5Y5b2hWisRnXXNUJOcPtkg1v2R9dZK"
	message := "eyJleHBpcmF0aW9uIjoiMjAyMy0xMi0wNVQwODowNzoyMS4xMjNaIiwiY29uZGl0aW9ucyI6W3siYnVja2V0IjoieGRsdW1pYTIifSx7ImtleSI6InVzZXItd2Vic2l0ZS1pY29uLzIwMjMxMjA0L1dOUTZPY19YcGNOdG11VGZxTk5INzA0ODUucG5nIn0sWyJjb250ZW50LWxlbmd0aC1yYW5nZSIsMCwxMDczNzQxODI0XV19"
	fmt.Println(b64_hmac(accessKeySecret, message))
}
