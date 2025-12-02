package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}

func RandomStr(len int) string {
	charset := "abcdefjhijklnmopqrstuvwxyzABCDEFJHIJKLNMOPQRSTUVWXYZ0123456789"
	str := make([]byte, len)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		str[i] = charset[r.Intn(62)]
	}
	return string(str)
}
