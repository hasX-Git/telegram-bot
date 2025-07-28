package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
	"time"
)

func RandWithRange(digits int) int {
	//digits is how much digits u want in a random number
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var min int = 1
	var max int = 9
	var i int
	for i = 0; i < digits-1; i++ {
		min = min * 10
		max = max*10 + 9
	}
	return r.Intn(max-min+1) + min
}

func CurrentDateAsID(n int) string {
	return strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(time.Now().Day()) + strconv.Itoa(RandWithRange(n))
}

func CheckValidityOfID(id string, n int) error {
	if len(id) != n {
		return errors.New("invalid ID")
	}

	//only digits
	for _, ch := range id {
		if int(ch) < 48 && int(ch) > 57 {
			return errors.New("invalid ID")
		}
	}
	return nil
}

func Hash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}
