package authx

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GetHashPassword(password string, secret string) string {
	sha := sha256.New()
	sha.Write([]byte(secret))
	sha.Write([]byte(password))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func GenerateRandomPassword(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[r.Intn(len(chars))])
	}
	return b.String()
}
