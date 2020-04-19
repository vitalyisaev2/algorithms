package utils

import (
	"math/rand"
)

func RandomBytes(n int) []byte {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return b
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomASCIIString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

var (
	Thai             = []int64{3585, 3654}
	Armenian         = []int64{1328, 1423}
	Chinese          = []int64{19968, 40869}
	JapaneseKatakana = []int64{12449, 12531}
	JapaneseHiragana = []int64{12353, 12435}
	KoreanHangul     = []int64{12593, 12686}
	CyrillicRussian  = []int64{1025, 1169}
	Greek            = []int64{884, 974}
)

func RandomUnicodeString(n int, charset []int64) string {
	rcv := make([]rune, n)

	for i := range rcv {
		rcv[i] = rune(charset[0] + rand.Int63n(charset[1]-charset[0]))
	}

	return string(rcv)
}

func ShuffleString(input string) string {
	runes := []rune(input)

	rand.Shuffle(len(runes), func(i, j int) { runes[i], runes[j] = runes[j], runes[i] })

	return string(runes)
}
