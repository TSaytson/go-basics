package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("https://google.com") // _ identifies an error not used
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", body)
}
