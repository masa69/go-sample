
package main


import (
	"fmt"
	"time"
)


func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(i + 1, "\n")
	}
}
