package main

import (
	"flag"
	"fmt"
	"github.com/eric-jxl/go_utils/utils"
)

func main() {
	i := flag.Int("i", 1, "Usage input Step number!")
	flag.Parse()
	s := utils.Stepbystep(*i)
	fmt.Printf("爬楼梯有%v种\n", s)
}
