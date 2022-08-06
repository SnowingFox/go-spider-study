package parser

import (
	"fmt"
	"math/rand"
	"spider/spider/engine"
	"spider/spider/model"
)

func generateSex() string {
	if rand.Intn(10) > 5 {
		return "man"
	} else {
		return "woman"
	}
}

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{
		Name: name,
		Age:  "19",
		Sex:  generateSex(),
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	fmt.Println(result)

	return result
}
