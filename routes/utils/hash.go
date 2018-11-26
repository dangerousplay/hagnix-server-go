package utils

import (
	"crypto/sha1"
	"fmt"
	"hash/fnv"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func HashStringSHA1(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return fmt.Sprintf("%s", h.Sum(nil))
}

func Hashcode(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func RandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func ToCommaSpaceSeparated(ints []int) string {
	var IDs []string
	for _, i := range ints {
		IDs = append(IDs, strconv.Itoa(i))
	}

	return strings.Join(IDs, ", ")
}

func FromCommaSpaceSeparated(csv string) ([]int, error) {
	var list []int

	for _, v := range strings.Split(csv, ", ") {
		inte, err := strconv.Atoi(v)
		if err != nil {
			return list, err
		}
		list = append(list, inte)
	}

	return list, nil
}

func AppendCommaSpaceSeparated(string string, ints []int) string {
	csv := ToCommaSpaceSeparated(ints)

	return string + ", " + csv
}
