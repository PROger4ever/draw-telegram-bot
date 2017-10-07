package common

import (
	"fmt"
	"os"
)

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

func PanicIfError(err error, message string) {
	if err != nil {
		panic(fmt.Errorf("Error occured while %v:\n%q", message, err))
	}
}

func RepairIfError(message string, data interface{}) {
	if r := recover(); r != nil {
		fmt.Fprint(os.Stderr, fmt.Errorf("Recovered while %v:\n%q\nThe data occured the panic: %q", message, r, data))
	}
}