package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//只读的方式
	file, err := os.Open("../struct_tag/main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read file:", err)
			return
		}
		fmt.Println(line)
	}
}
