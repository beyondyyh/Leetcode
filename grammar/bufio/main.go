package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	in, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.OpenFile("2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	inReader := bufio.NewReader(in)
	outWriter := bufio.NewWriter(out)
	for {
		line, err := inReader.ReadString('\n')
		log.Printf("read line:%s err:%+v\n", strings.TrimRight(line, "\n"), err)
		if err == io.EOF {
			break
		}

		n, err := outWriter.WriteString(line)
		log.Printf("write len:%d err:%+v\n", n, err)
	}

	outWriter.WriteString("the end")
	outWriter.Flush()
}
