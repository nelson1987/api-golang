package main

import (
	"fmt"
	"net/http"
)

func main() {
	site := "http://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}
