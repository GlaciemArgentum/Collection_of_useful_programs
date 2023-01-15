package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Properties struct {
	Id int `json:"global_id"`
}

func CheckErr(err error, s string) {
	if err != nil {
		str := fmt.Sprintf("%s: %v\n", s, err)
		panic(str)
	}
}

func main() {
	file, err := os.Open("./data/data-20190514T0100.json")
	CheckErr(err, "Ошибка открытия файла")
	defer func() {
		err := file.Close()
		CheckErr(err, "Ошибка закрытия файла!")
	}()

	data, err := io.ReadAll(file)
	CheckErr(err, "Ошибка чтения файла")

	var decode []Properties
	err = json.Unmarshal(data, &decode)
	CheckErr(err, "Ошибка раскодирования файла")

	sum := 0
	for _, i := range decode {
		sum += i.Id
	}

	fmt.Println(sum)
}
