package main

import (
	"math/rand"
	"time"
)

func GetFixedHash(len int) string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, len)
	for i := range b {
		b[i] = letterBytes[rand.Intn(62)]
	}
	return string(b)
}
