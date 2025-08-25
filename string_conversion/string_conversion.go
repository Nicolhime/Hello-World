package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(printer(123))

}

func printer(x int) string {

	var y string = strconv.Itoa(x)

	return y
}
