package fetch

import (
	"fmt"
	"io"
	"net/http"
)

var baseUrl = "https://www.zhenai.com/zhenghun"

func Fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}

	return io.ReadAll(res.Body)
}
