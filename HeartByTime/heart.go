package main
//test label
import (
	"fmt"
	"time"
)

func main() {

	for y := 1.5; y > -1.5; y -= 0.1 {

		//fmt.Println(y)
		for x := -1.5; x < 1.5; x += 0.05 {
			//fmt.Println(x)
			a := x*x + y*y - 1
			//fmt.Println(a)

			if a*a*a-x*x*y*y*y <= 0.0 {
				fmt.Printf("*")
			} else {
				fmt.Printf("_")
			}

		}

		time.Sleep(time.Millisecond * 100)
		fmt.Println(" ")

	}
}
