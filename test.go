package printer

import (
	"errors"
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
		err      error
	}{
		{[]byte("Hello, World!"), 13, nil},            // Простой ASCII текст
		{[]byte("Привет, мир!"), 13, nil},             // Русский текст
		{[]byte("こんにちは"), 5, nil},                     // Японский текст
		{[]byte("👋🌍"), 2, nil},                        // Эмодзи
		{[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8}, // Некорректный UTF-8
		{[]byte{0xe0, 0xa4}, 0, ErrInvalidUTF8},       // Некорректный UTF-8 (недостаточно байт)
		{[]byte(""), 0, nil},                          // Пустая строка
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			length, err := GetUTFLength(test.input)
			if length != test.expected || !errors.Is(err, test.err) {
				t.Errorf("expected (%d, %v), got (%d, %v)", test.expected, test.err, length, err)
			}
		})
	}
}
