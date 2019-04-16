package helper

import "math/rand"

//RandomStrings returns 随机的字符串
func RandomStrings(size int) string {
	source := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sourceLen := len(source)

	s := ""
	for i := 0; i < size; i++ {
		n := rand.Intn(sourceLen)
		s += source[n : n+1]
	}

	return s
}
