package main

import (
	"fmt"
)

func main() {
	fmt.Println(`The value of "true && true" is`, true)
	fmt.Println(`The value of "true && false" is`, true && false)
	fmt.Println(`The value of "true || true" is`, true)
	fmt.Println(`The value of "true || false" is`, true || false)
	fmt.Println(`The value of "!true" is`, !true)
}
