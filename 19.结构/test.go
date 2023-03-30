package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Name string  `json:"name"`
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func main() {
	loc := []location{
		{Name: "A", Lat: 32.8771, Long: 162.32341},
		{Name: "B", Lat: 32.8771, Long: 162.32341},
		{Name: "C", Lat: 32.8771, Long: 162.32341},
	}

	bytes, err := json.Marshal(loc)
	exitOnError(err)

	fmt.Println(string(bytes))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
