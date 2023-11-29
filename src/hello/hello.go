package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(20 * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println(resp)
	} else {
		fmt.Println(err)
	}
}
