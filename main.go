package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

var data ColorGroup

func main() {
	Marshal()
	Unmarshal()
}
func Marshal() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}
func Unmarshal() {
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	json.Unmarshal(val, &data)
}
