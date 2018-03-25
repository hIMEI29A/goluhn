# goluhn

[![GoDoc](https://godoc.org/github.com/hIMEI29A/golun?status.svg)](http://godoc.org/github.com/hIMEI29A/goluhn)

**goluhn** - Golang implementation of Luhn algorithm as it described (on russian) [here](https://ru.wikipedia.org/wiki/%D0%90%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC_%D0%9B%D1%83%D0%BD%D0%B0#%D0%A3%D0%BF%D1%80%D0%BE%D1%89%D1%91%D0%BD%D0%BD%D1%8B%D0%B9_%D0%B0%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC)

## Description

**Goluhn** exports 3 functions. This functions is

```go
// CheckLuhn returns true if Lun checksum of given value is correct
func CheckLuhn(num int) bool {}

// LuhnInRange returns a slice of integers with correct Lun checksum in given range
func LuhnInRange(start, end int) []int {}

// LuhnByLen returns random int with correct Lun checksum and given length
func LuhnByLen(len int) int {}
```

## Install

	go get github.com/hIMEI29A/goluhn
