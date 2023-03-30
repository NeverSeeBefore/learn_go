package main

import "fmt"

type planets []string

func (ps planets) terraform() {
	for i := range ps {
		ps[i] = "New " + ps[i]
	}
}

func main() {
	ps := []string{"Mars", "Uranus", "Neptune"}

	planets.terraform(ps)

	fmt.Println(ps)

}
