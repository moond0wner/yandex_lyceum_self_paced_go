// "Тест GetUTFLength"

package main

import (
	"errors"
	"testing"
)

type Test struct {
	input          []byte
	expectedLength int
	expectedError  error
}

var tests = []Test{
	{[]byte("Hello, World!"), 13, nil},            // Простой текст
	{[]byte("こんにちは"), 5, nil},                     // Японский текст
	{[]byte("😀"), 1, nil},                         // Эмодзи
	{[]byte{0xff, 0xff, 0xff}, 0, ErrInvalidUTF8}, // Неверная UTF-8
	{[]byte("Hello, 世界"), 9, nil},                 // Комбинация символов
	{[]byte{0xE2, 0x82, 0xAC}, 1, nil},            // Корректный UTF-8 (евро)
	{[]byte(""), 0, nil},                          // Пустая строка
	{[]byte{0xC3, 0x28}, 0, ErrInvalidUTF8},       // Неверная последовательность
}

func TestGetUTFLength(t *testing.T) {
	for _, test := range tests {
		length, err := GetUTFLength(test.input)
		if length != test.expectedLength || !errors.Is(err, test.expectedError) {
			t.Fatalf(`GetUTFLength(%q)=(%d, %v); expected (%d, %v)`, test.input, length, err, test.expectedLength, test.expectedError)
		}
	}
}
