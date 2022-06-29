package main

import (
	"fmt"
	"flag"
 )
func stepbystep(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return stepbystep(n-1) + stepbystep(n-2)
}

func main() {
	i := flag.Int("i",1,"Usage input Step number!")
	flag.Parse()
	s :=stepbystep(*i)
	fmt.Printf("爬楼梯有%v种\n",s)
}
