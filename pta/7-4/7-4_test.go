package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestASCII(t *testing.T) {
	tests := []struct {
		name string
		str  rune
		want int
	}{
		{"case 1", 'A', 65},
		{"case 2", '0', 48},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, int(tt.str))
	}
}
