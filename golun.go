package golun

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// ErrFatal is a basic error handler
func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Chod simply check is number is even or odd. It returns true if even
func chod(num int) bool {
	// odd
	check := false
	if num%2 == 0 {
		// even
		check = true
	}

	return check
}

// Powt returns int converted 10**p
func powt(p int) int {
	return int(math.Pow10(p))
}

func randt(num int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	random := r.Intn(num)

	return random
}

// AToi converts string to int
func aToi(str string) int {
	atoi, err := strconv.Atoi(str)
	errFatal(err)

	return atoi
}

// Base10 simply checks if given num is divisible by 10
func base10(num int) bool {
	var check bool

	if num%10 == 0 {
		check = true
	} else {
		check = false
	}

	return check
}

// Base9 makes Lun check's performans with given int
func base9(num int) int {
	var check int

	if 2*num > 9 {
		check = (2 * num) - 9
	} else {
		check = 2 * num
	}

	return check
}

// CheckLuhn checks is Lun checksum of given value is correct
func CheckLuhn(num int) bool {
	var check int

	numstring := strconv.Itoa(num)

	switch {
	case chod(len(numstring)) == true:
		for i, n := range numstring {
			if chod(i+1) == false {
				check += base9(aToi(string(n)))
			} else {
				check += aToi(string(n))
			}
		}

	case chod(len(numstring)) != true:
		for i, n := range numstring {
			if chod(i+1) == true {
				check += base9(aToi(string(n)))
			} else {
				check += aToi(string(n))
			}
		}
	}

	return base10(check)
}

// LuhnInRange returns a slice of integers with correct Lun checksum in given range
func LuhnInRange(start, end int) []int {
	var nums []int

	for i := start; i <= end; i++ {
		if CheckLuhn(i) == true {
			nums = append(nums, i)
		}
	}

	return nums
}

// LuhnByLen returns random int with correct Lun checksum and given length
func LuhnByLen(len int) int {
	var luhn int
	base := powt(len - 1)
	rangenums := (powt(len) - 1) - powt(len-1)

	preluhn := base + randt(rangenums)

	switch {
	case CheckLuhn(preluhn):
		luhn = preluhn

	default:
		luhn = LuhnByLen(len)
	}

	return luhn
}
