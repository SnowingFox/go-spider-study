package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var baseUrl = "https://www.zhenai.com/zhenghun"

func main() {
	res, err := http.Get(baseUrl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error code", res.StatusCode)
		return
	}

	all, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	printCityList(all)
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {

		fmt.Printf("%s %s", m[1], m[2])
	}
}
