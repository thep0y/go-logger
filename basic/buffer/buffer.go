/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: buffer.go
 * @Created: 2021-05-16 09:06:54
 * @Modified: 2021-05-16 09:19:52
 */

package buffer

type Buffer []byte

// Reset 重置缓冲
func (b *Buffer) Reset() {
	*b = Buffer([]byte(*b)[:0])
}

// Append 添加 byte 切片到缓冲
func (b *Buffer) Append(data []byte) {
	*b = append(*b, data...)
}

// AppendByte 添加 byte 到缓冲
func (b *Buffer) AppendByte(data byte) {
	*b = append(*b, data)
}

// AppendInt 添加 int 到缓冲
func (b *Buffer) AppendInt(val int, width int) {
	var repr [8]byte
	reprCount := len(repr) - 1
	for val >= 10 || width > 1 {
		reminder := val / 10
		repr[reprCount] = byte('0' + val - reminder*10)
		val = reminder
		reprCount--
		width--
	}

	repr[reprCount] = byte('0' + val)
	b.Append(repr[reprCount:])
}

func (b Buffer) Bytes() []byte {
	return []byte(b)
}
