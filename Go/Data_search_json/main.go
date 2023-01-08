package main

import "os"

type Properties struct {
	Id int `json:"global_id"`
}

func main() {
	file, err := os.Open("./data/")
}
