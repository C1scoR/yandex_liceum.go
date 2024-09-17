package main

import (
    "errors"
    "testing"
)

func TestGetUTFLength(t *testing.T) {
    tests := []struct {
        name     string
        input    []byte
        expected int
        err      error
    }{
        {"Valid UTF8 string", []byte("Hello"), 5, nil},
        {"Valid UTF8 with multibyte", []byte("Привет"), 6, nil},
        {"Empty string", []byte(""), 0, nil},
        {"Invalid UTF8", []byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},
        {"Valid UTF8 emoji", []byte("😊"), 1, nil},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := GetUTFLength(tt.input)
            
            if result != tt.expected {
                t.Errorf("GetUTFLength(%q) = %d; want %d", tt.input, result, tt.expected)
            }

            if err != tt.err {
                t.Errorf("GetUTFLength(%q) error = %v; want %v", tt.input, err, tt.err)
            }
        })
    }
}