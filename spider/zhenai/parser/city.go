package parser

import (
	"regexp"
	"spider/spider/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: "http://www.zhenai.com/zhenghun/anhui",
			ParserFunc: func(c []byte) engine.ParseResult {
				ParseProfile(c, name)

				return engine.ParseResult{}
			},
		})

	}

	return result
}
