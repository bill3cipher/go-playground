package main

import (
	"errors"
	"fmt"
)

// 1. ใช้ errors.New ในการสร้าง error
// 2. type ErrTooYoung int func (err ErrTooYoung) Error() string { return "" }

type ErrTooYoung int

func (err ErrTooYoung) Error() string {
	return fmt.Sprintf("%d is too young", err)
}

// validateLength return false when string length less than 8
// Please change return type to error with message "too short string"
func validateLength(s string) error {
	if len([]rune(s)) < 8 {
		return errors.New("too short string")
	}
	return nil
}

var ageError = errors.New("your age is negative!")

// in case of too young age please create a new type ErrTooYoung int` with method `Error() string`
// example error message: "17 is too young"
func validateAge(n int) error {
	if n < 0 {
		return ageError
	}
	if n < 18 {
		return ErrTooYoung(n)
	}
	return nil
}

func main() {
	errLength := validateLength("a")
	fmt.Println(errLength)
	errLength2 := validateLength("aaaaaaaaaaaaaa")
	fmt.Println(errLength2)
	errAge := validateAge(17)
	fmt.Println(errAge)
	errAge2 := validateAge(-1)
	fmt.Println(errAge2)
}
