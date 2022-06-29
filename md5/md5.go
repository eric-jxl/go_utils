package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
)


func main() {
	md5_str := md5.New()
	s := flag.String("s", "123456", "Usage String to MD5")
	flag.Parse()
	md5_str.Write([]byte(*s))
	hex_md5_str := hex.EncodeToString(md5_str.Sum(nil))
	fmt.Println(hex_md5_str)
}

