package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

type HexChunk string

type HexChunks []HexChunk

type encodingTable map[rune]string

const chunkSize = 8

func Encode(str string) string {
	str = prepareText(str)

	chunks := splitByChunks(encodeBin(str), chunkSize)

	return chunks.ToHex().ToString()
}

func (hcs HexChunks) ToString() string {
	const sep = " "

	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	}

	var b strings.Builder

	b.WriteString(string(hcs[0]))

	for _, hc := range hcs[1:] {
		b.WriteString(sep)
		b.WriteString(string(hc))
	}

	return b.String()
}

func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		hexChunk := chunk.ToHex()

		res = append(res, hexChunk)
	}

	return res
}

func (bcs BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bcs), 2, chunkSize)
	if err != nil {
		panic(err)
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

func splitByChunks(str string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(str) / chunkSize

	chunksCount := strLen / chunkSize

	if strLen%chunkSize != 0 {
		chunksCount++
	}

	result := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, c := range str {
		buf.WriteString(string(c))

		if (i+1)%chunkSize == 0 {
			result = append(result, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastchunk := buf.String()

		lastchunk += strings.Repeat("0", chunkSize-len(lastchunk))

		result = append(result, BinaryChunk(lastchunk))

	}

	return result
}

func encodeBin(str string) string {
	var buffer strings.Builder

	for _, char := range str {
		buffer.WriteString(bin(char))
	}
	return buffer.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown char: " + string(ch))
	}

	return res
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
