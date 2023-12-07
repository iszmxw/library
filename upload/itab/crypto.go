package itab

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

func b64_hmac(str1, str2 string) string {
	return base64Encode(encryptOrHash(str1, str2), byte('='))
}

// base64Encode 将输入字符串进行 Base64 编码
func base64Encode(input string, paddingChar byte) string {
	// 如果未提供填充字符，默认使用 '='
	if paddingChar == 0 {
		paddingChar = '='
	}

	var result []byte
	inputLength := len(input)

	for i := 0; i < inputLength; i += 3 {
		// 将三个字符合并成一个数字
		combined := int(input[i])<<16 | func() int {
			if i+1 < inputLength {
				return int(input[i+1]) << 8
			}
			return 0
		}() | func() int {
			if i+2 < inputLength {
				return int(input[i+2])
			}
			return 0
		}()

		for j := 0; j < 4; j++ {
			// 判断是否需要填充
			if 8*i+6*j > 8*inputLength {
				result = append(result, paddingChar)
			} else {
				// 获取 Base64 字符并拼接到结果中
				result = append(result, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"[(combined>>(6*(3-j)))&63])
			}
		}
	}

	return string(result)
}

/**
 * 对输入字符串进行加密或哈希处理
 * @param {string} str1 - 第一个输入字符串
 * @param {string} str2 - 第二个输入字符串
 * @returns {string} - 处理后的字符串结果
 */
func encryptOrHash(str1, str2 string) string {
	// 如果条件为真，则对输入字符串进行 URI 编码和解码
	str1 = uriEncodeDecode(str1)
	str2 = uriEncodeDecode(str2)

	// 将 UTF-8 编码的字符串转换为字节数组
	bytes1 := utf8ToByteArray(str1)

	// 如果字节数组长度大于16，则进行填充操作
	if len(bytes1) > 16 {
		bytes1 = padInput(bytes1, 8*len(str1))
	}

	// 初始化两个数组
	var xorArray1 [16]int32
	var xorArray2 [16]int32

	// 对数组进行异或运算
	for i := 0; i < len(xorArray2); i++ {
		if i < len(bytes1) {
			bytesRes := bytes1[i]
			xorArray1[i] = 909522486 ^ bytesRes
			xorArray2[i] = 1549556828 ^ bytesRes
		} else {
			xorArray1[i] = 909522486
			xorArray2[i] = 1549556828
		}
	}

	// 将第二个字符串转换为字节数组，并与填充后的第一个数组拼接
	concat := append(xorArray1[:], utf8ToByteArray(str2)...)
	str2Len := 512 + 8*len(str2)
	fmt.Println(str2Len, concat)
	bytes2 := padInput(concat, str2Len)
	fmt.Println(bytes2)
	fmt.Println(bytes2)

	// 将两个数组拼接并填充到672位，然后转换为字符串返回结果
	return convertArrayToString(padInput(append(xorArray2[:], bytes2...), 672))
}

// Utf8ToByteArray 将字符串转换为 UTF-8 编码的字节数组
func utf8ToByteArray(input string) []int32 {
	bitLength := utf8.RuneCountInString(input) * 8
	lenCap := utf8.RuneCountInString(input) >> 2
	byteArray := make([]int32, lenCap)
	// 初始化字节数组
	for i := 0; i < lenCap; i++ {
		byteArray[i] = 0
	}
	for j := 0; j < bitLength; j += 8 {
		charCodeAt := []rune(input)[j/8]
		if j>>5 > len(byteArray)-1 {
			continue
		}
		byteArray[j>>5] |= (255 & int32(charCodeAt)) << (24 - j%32)
	}
	return byteArray
}

// 对字符串进行两次 URI 编码和解码，用于确保特殊字符正确处理
func uriEncodeDecode(str string) string {
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		panic(err)
	}
	decodedStr := encodeURIComponent(unescape)
	return decodedStr
}

// 可以通过修改底层url.QueryEscape代码获得更高的效率，很简单
func encodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	r = strings.Replace(r, "%21", "!", -1)
	return r
}

func removeZeros(slice []uint32) []uint32 {
	var result []uint32

	for _, value := range slice {
		if value != 0 {
			result = append(result, value)
		}
	}

	return result
}

// 执行 SHA-1 哈希算法的填充操作
func padInput(message []int32, length int) []int32 {
	w := make([]int32, 80)
	h0 := int32(1732584193)
	h1 := int32(-271733879)
	h2 := int32(-1732584194)
	h3 := int32(271733878)
	h4 := int32(-1009589776)
	// 进行填充
	index := 15 + (length+64)>>9<<4
	key := length >> 5
	if key <= len(message)-1 {
		message[key] |= 128 << (24 - length%32)
	}
	// 直接添加元素到索引 4 的位置
	indexToAdd := index
	if indexToAdd <= len(message)-1 {
		// 如果索引在当前切片的范围内，直接设置值
		message[indexToAdd] = int32(uint32(length))
	}
	for block := 0; block < len(message); block += 16 {
		// 初始化哈希值
		a := h0
		b := h1
		c := h2
		d := h3
		e := h4
		// 处理每个 512 位块
		for i := 0; i < 80; i++ {
			if i < 16 {
				if block+i <= len(message)-1 {
					w[i] = message[block+i]
				}
			} else {
				w[i] = leftRotate(w[i-3]^w[i-8]^w[i-14]^w[i-16], 1)
			}

			temp := safeAdd(safeAdd(leftRotate(a, 5), int32(sha1F(int32(i), b, c, d))), safeAdd(safeAdd(e, w[i]), sha1K(int32(i))))
			e = d
			d = c
			c = leftRotate(b, 30)
			b = a
			a = temp
		}
		// 更新哈希值
		h0 = safeAdd(h0, a)
		h1 = safeAdd(h1, b)
		h2 = safeAdd(h2, c)
		h3 = safeAdd(h3, d)
		h4 = safeAdd(h4, e)
	}
	// 返回最终的哈希值数组
	return []int32{h0, h1, h2, h3, h4}
}

// SHA-1 算法中使用的常量 K
func sha1K(t int32) int32 {
	switch {
	case t < 20:
		return 1518500249
	case t < 40:
		return 1859775393
	case t < 60:
		return -1894007588
	default:
		return -899497514
	}
}

// SHA-1 算法中使用的函数 F
func sha1F(t, b, c, d int32) int32 {
	if t < 20 {
		return (b & c) | (^b & d) // 0 <= t < 20
	} else if t < 40 {
		return b ^ c ^ d // 20 <= t < 40
	} else if t < 60 {
		return (b & c) | (b & d) | (c & d) // 40 <= t < 60
	} else {
		return b ^ c ^ d // 60 <= t < 80
	}
}

// 辅助函数：左移
func leftRotate(value, shift int32) int32 {
	return (value << shift) | (value >> (32 - shift))
}

// 辅助函数：安全的加法，处理位溢出
func safeAdd(x, y int32) int32 {
	// 获取低16位的和
	lsw := (x & 0xFFFF) + (y & 0xFFFF)
	// 获取高16位的和，并考虑低位的进位
	msw := (x >> 16) + (y >> 16) + (lsw >> 16)
	// 将结果组合成32位整数
	return (msw << 16) | (lsw & 0xFFFF)
}

func convertArrayToString(array []int32) string {
	// 计算字符串的长度，每个元素占32位
	length := 32 * len(array)
	result := ""

	// 迭代每8位，从数组中提取相应的字节，并构建字符串
	for i := 0; i < length; i += 8 {
		// 获取当前字节的字符编码
		charCode := (array[i>>5] >> uint(24-i%32)) & 255
		// 将字符编码转换为字符串并添加到结果字符串
		result += string(rune(charCode))
	}

	// 返回最终的字符串
	return result
}
