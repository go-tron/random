package random

import (
	"math/rand"
)

const numberSource = "0123456789"
const stringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const letterSource = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Number(n int) string {
	b := make([]byte, n)
	for i := range b {
		rd := rand.Intn(len(numberSource))
		b[i] = numberSource[rd]
	}
	return string(b)
}

func String(n int) string {
	b := make([]byte, n)
	for i := range b {
		rd := rand.Intn(len(stringSource))
		b[i] = stringSource[rd]
	}
	return string(b)
}

func Letter(n int) string {
	b := make([]byte, n)
	for i := range b {
		rd := rand.Intn(len(letterSource))
		b[i] = letterSource[rd]
	}
	return string(b)
}

func Source(n int, source string) string {
	b := make([]byte, n)
	for i := range b {
		rd := rand.Intn(len(source))
		b[i] = source[rd]
	}
	return string(b)
}
