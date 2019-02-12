package utils

import (
	"time"
	"fmt"
	"github.com/badoux/checkmail"
)

func DateTimeFormat() string {
	now := time.Now()
	var day, month, year int = now.Day(), int(now.Month()), now.Year()
	var hour, min, sec int = now.Hour(), now.Minute(), now.Second()
	return fmt.Sprintf("%d-%d-%d T%d:%d:%d", day, month, year, hour, min, sec)
}

func IsEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}