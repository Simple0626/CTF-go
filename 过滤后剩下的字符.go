package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `[b-zA-Z_@#%^&*:{}\-\+<>\"|` + "`" + `\;\[\]]`
	//println(str)
	var left_str string
	re := regexp.MustCompile(str)
	for i := 32; i < 128; i++ {
		//fmt.Println(string(i))
		//if string(i) == "/" {
		//	fmt.Println(i)
		//}
		s := re.FindAllString(string(i), -1)
		fmt.Println(s)
		if len(s) == 0 {
			left_str += string(i)
		}

	}
	fmt.Printf("剩下的字符：%s", left_str)

}
