package main

import (
	"regexp"
	"fmt"
	"os"
)

/*
三个函数来判断是否匹配：
Match
MatchReader
MatchString
*/
func main() {
	fmt.Println(IsIp("1.1.1.1"))

	//判断输入字符串
	if len(os.Args) == 1 {
		fmt.Println()
	}else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {

	}

	//复杂正则
	s := `123456abcd""`
	re, _ := regexp.Compile("^[0-9]+")
	s = re.ReplaceAllString(s, "a")
	fmt.Println(s)
}

//判断ip
func IsIp(ip string) (b bool) {
	if m,_ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}
