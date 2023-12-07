package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"net/url"
)

func b64_hmac(t, e string) string {
	return customBase64Encode(string(d([]byte(t), []byte(e))), byte('='))
}

func customBase64Encode(input string, paddingChar byte) string {
	var result string
	inputBytes := []byte(input)

	for i := 0; i < len(inputBytes); i += 3 {
		val := int(inputBytes[i])<<16 | getCharCodeAt(inputBytes, i+1)<<8 | getCharCodeAt(inputBytes, i+2)

		for j := 0; j < 4; j++ {
			if 8*i+6*j > 8*len(inputBytes) {
				result += string(paddingChar)
			} else {
				result += string("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"[val>>6*(3-j)&63])
			}
		}
	}

	return result
}

func getCharCodeAt(input []byte, index int) int {
	if index < len(input) {
		return int(input[index])
	}
	return 0
}

func uint32toByte(uint32Slice []uint32) []byte {
	// 使用 bytes.Buffer 保存二进制数据
	var buf bytes.Buffer

	// 将 []uint32 写入到 Buffer 中
	err := binary.Write(&buf, binary.LittleEndian, uint32Slice)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return nil
	}

	// 从 Buffer 中读取 []byte
	return buf.Bytes()
}
func d(t, e []byte) []byte {
	var n, o, s, a, u uint32

	t = r(t)
	e = r(e)

	n = binary.LittleEndian.Uint32(uint32toByte(c(t)[:4]))

	o = binary.LittleEndian.Uint32(uint32toByte(c(t)[4:8]))
	s = binary.LittleEndian.Uint32(uint32toByte(c(t)[8:12]))
	a = binary.LittleEndian.Uint32(uint32toByte(c(t)[12:16]))

	if len(c(t)) > 16 {
		hashed := sha256.Sum256(t)
		n = binary.LittleEndian.Uint32(hashed[:4])
		o = binary.LittleEndian.Uint32(hashed[4:8])
		s = binary.LittleEndian.Uint32(hashed[8:12])
		a = binary.LittleEndian.Uint32(hashed[12:16])
	}

	var oArray, sArray [16]uint32

	for i := 0; i < 16; i++ {
		oArray[i] = 909522486 ^ n
		sArray[i] = 1549556828 ^ n
	}

	concatenated := append(oArray[:], c(e)...)
	u = binary.LittleEndian.Uint32(uint32toByte(p(concatenated, 0)[:4]))

	hashed := sha256.Sum256(append(sArray[:], uint32toByte(p(concatenated, 0)...)))
	return p(append(sArray[:], hashed[:]...))
}

func c(t []byte) []uint32 {
	var e, n = 8 * len(t), make([]uint32, (len(t)+3)>>2)

	for i := 0; i < len(n); i++ {
		n[i] = 0
	}

	for i := 0; i < e; i += 8 {
		n[i>>5] |= (uint32(t[i/8]) & 255) << (24 - uint32(i%32))
	}

	return n
}

func r(t []byte) []byte {
	// Use url.QueryEscape to replace JavaScript's encodeURIComponent
	encoded := url.QueryEscape(string(t))
	// Use url.PathEscape if you want to mimic encodeURIComponent more closely
	// encoded := url.PathEscape(string(t))

	// Convert the result back to []byte
	return []byte(encoded)
}

func p(t []uint32, e int) []uint32 {
	var n, r, a, i, c, u, h, f []uint32
	p := make([]uint32, 80)
	l := uint32(1732584193)
	d := uint32(-271733879)
	v := uint32(-1732584194)
	b := uint32(271733878)
	m := uint32(-1009589776)

	t[e>>5] |= 128 << (24 - e%32)
	t[15+(e+64>>9<<4)] = uint32(e)

	for n = 0; n < len(t); n += 16 {
		i = l
		c = d
		u = v
		h = b
		f = m
		r = make([]uint32, 80)

		for a := 0; a < 80; a++ {
			if a < 16 {
				r[a] = t[n+a]
			} else {
				r[a] = s(r[a-3]^r[a-8]^r[a-14]^r[a-16], 1)
			}
			i, c, u, h, f = o(i, c, u, h, f, r[a])
		}

		l, d, v, b, m = o(l, d, v, b, m, i, c, u, h, f)
	}
	return []uint32{l, d, v, b, m}
}

func o(t, e, n, r, a, i uint32) (uint32, uint32, uint32, uint32, uint32) {
	var c uint32
	c = (t + (e&n | ^e&r) + a + i) << 32
	return (t >> 32) + (e >> 32) + (n >> 32) + (r >> 32) + (c >> 32) + (t & 0xFFFFFFFF) + (e & 0xFFFFFFFF) + (n & 0xFFFFFFFF) + (r & 0xFFFFFFFF) + (c & 0xFFFFFFFF)
}

func i(t []uint32) []byte {
	var e, n int
	var r []byte
	for e = 0; e < 32*len(t); e += 8 {
		n = e >> 5
		r = append(r, byte(t[n]>>(24-e%32)&255))
	}
	return r
}

func g(t, e, n, r uint32) uint32 {
	switch {
	case t < 20:
		return (e & n) | (^e & r)
	case t < 40:
		return e ^ n ^ r
	case t < 60:
		return (e & n) | (e & r) | (n & r)
	default:
		return e ^ n ^ r
	}
}

func s(t, e uint32) uint32 {
	return (t << e) | (t >> (32 - e))
}

func y(t int) uint32 {
	switch {
	case t < 20:
		return 1518500249
	case t < 40:
		return 1859775393
	case t < 60:
		return 0xF0E65480
	default:
		return 0xC3D2E1F0
	}
}

func main() {
	// 在这里调用 hashFunction 方法并打印结果
	result := b64_hmac("your_message", "your_key")
	fmt.Println(result)
}
