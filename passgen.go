package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return (string(b))
}

func main() {
	serverFlag := flag.Int("l", int(16), "lenght, a number")

	flag.Parse()

	leng := *serverFlag

	fmt.Println(RandStringRunes(leng))
}
