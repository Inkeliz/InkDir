package Display

import (
	"encoding/binary"
)

var (
	w1_MIN = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x00})
	w1_MAX = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x19})
	w2_MIN = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x20})
	w2_MAX = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x1F, 0xFF})
	w3_MIN = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x20, 0x00})
	w3_MAX = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0xFF, 0x60})
	w4_MIN = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0xFF, 0x61})
	w4_MAX = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0xFF, 0x9F})
	w5_MIN = binary.BigEndian.Uint32([]byte{0x00, 0x00, 0xFF, 0xA0})
)

// Improved `str_pad` from PHP:
func StrpadRight(text string, pad string, padLength int) string {
	textWidth := Strwidth(text)
	padWidth := Strwidth(pad)

	for textWidth < padLength {
		text = text + pad
		textWidth += padWidth
	}
	return text
}

// Improved `str_pad` from PHP:
func StrpadLeft(text string, pad string, padLength int) string {
	textWidth := Strwidth(text)
	padWidth := Strwidth(pad)

	for textWidth < padLength {
		text = pad + text
		textWidth += padWidth
	}
	return text
}

// Based on `mb_strimwidth` from PHP:
func Strimwidth(text string, width int, trimmaker string) string {
	textWidth := Strwidth(text)

	switch {
	case width > textWidth:
		return text
	case width == 0:
		return trimmaker
	}
	return text[:width] + trimmaker
}

// Based on `mb_strwidth` from PHP:
func Strwidth(text string) int {
	width := 0

	for _, letter := range text {
		li := uint32(letter)

		switch {
		case li >= w1_MIN && li <= w1_MAX:
			width += 0
		case li >= w2_MIN && li <= w2_MAX:
			width += 1
		case li >= w3_MIN && li <= w3_MAX:
			width += 2
		case li >= w4_MIN && li <= w4_MAX:
			width += 1
		case li >= w5_MIN:
			width += 2
		}
	}
	return width
}
