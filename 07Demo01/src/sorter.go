package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"time"
	"algorithms/bubblesort"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a","qsort", "Sort algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("读取：", values)
		t1 := time.Now()
		switch *algorithm {
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("没有找到算法")
		}
		t2 := time.Now()
		fmt.Println("开始:",t1,"结束:",t2)
		writeValues(values, *outfile)
	}else {
		fmt.Println(err)
	}
}

//读文件
func readValues(infile string)(values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("打开文件失败", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("太长了")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

//写文件
func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("创建文件失败", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str+"\n")
	}
	return nil
}
