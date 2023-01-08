package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("task.data")
	if err != nil {
		str := fmt.Sprintf("Ошибка открытия файла: %v\n", err)
		panic(str)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			str := fmt.Sprintf("Ошибка закрытия файла: %v\n", err)
			panic(str)
		}
	}()

	r := bufio.NewReader(file)

	for i := 1; ; i++ {
		s, err := r.ReadString(';')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Конец файла")
				break
			}
			str := fmt.Sprintf("Ошибка считывания из файла: %v\n", err)
			panic(str)
		}
		s = strings.Replace(s, ";", "", -1)

		if s == "0" {
			fmt.Println(i)
			break
		}
	}
}
