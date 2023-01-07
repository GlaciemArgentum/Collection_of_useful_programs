package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Ой-ой, файл не открывается:", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Ошибка закрытия файла:", err)
		}
	}()

	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		fmt.Println("Ошибка считывания файла:", err)
	}
	if l := len(data); l == 10 {
		fmt.Printf("Answer: %v\n", data[4][2])
	}
	return nil
}

func main() {
	const root = "./task"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Println("Ошибка:", err)
	}
}
