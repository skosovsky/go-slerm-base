package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func itoa(a int) string {
	return strconv.Itoa(a)
}

func sprint(a int) string {
	return fmt.Sprint(a) //nolint:perfsprint // it's test
}

func itoaPlus(a int) string {
	return "" + strconv.Itoa(a)
}

func appendInt(a int) string {
	return string(strconv.AppendInt([]byte{}, int64(a), 10)) //nolint:gomnd // it's learning code
}
