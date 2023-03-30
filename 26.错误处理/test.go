package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	url1, err := url.Parse("a1b2bfn4mqkm://eqwreq 42452 ?aeq 01")

	if err != nil {
		if errs, ok := err.(*url.Error); ok {
			fmt.Printf("url1 错误断言 \n Err: %#v \n Op: %#v \n URL: %#v\n", errs.Err, errs.Op, errs.URL)
		}
		fmt.Printf("url1 error: %v\n", err)
	}
	fmt.Printf("%#v\n", url1)
	fmt.Println()

	u, err := url.Parse("http://example.com/#x/y%2Fz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fragment:", u.Fragment)
	fmt.Println("RawFragment:", u.RawFragment)
	fmt.Println("EscapedFragment:", u.EscapedFragment())

}
