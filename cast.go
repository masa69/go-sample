
package main


import (
	"fmt"
	"strconv"
)


type Cast struct {}


func NewCast() *Cast {
	return new(Cast)
}


func (c *Cast) strToBool(str string) {
	fmt.Print(strconv.ParseBool(str))
	fmt.Printf("\n")
}


func (c *Cast)  Atoi(str string) {
	fmt.Print(strconv.Atoi(str))
	fmt.Printf("\n")
}


func main() {
	cast := NewCast()

	fmt.Printf("[string to bool]\n")
	cast.strToBool("0")// false <nil>
	cast.strToBool("1")// true <nil>
	cast.strToBool("2")// false strconv.ParseBool: parsing "2": invalid syntax
	cast.strToBool("a")// false strconv.ParseBool: parsing "a": invalid syntax
	cast.strToBool("01")// false strconv.ParseBool: parsing "01": invalid syntax
	cast.strToBool("-1")// false strconv.ParseBool: parsing "-1": invalid syntax

	fmt.Printf("[string to int]\n")
	cast.Atoi("0")// 0 <nil>
	cast.Atoi("1.0")// 0 strconv.Atoi: parsing "1.0": invalid syntax
	cast.Atoi("a")// 0 strconv.Atoi: parsing "a": invalid syntax
	cast.Atoi("01")// 1 <nil>
	cast.Atoi("-1")// -1 <nil>
}