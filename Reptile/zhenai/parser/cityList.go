package parser

import (
	"GoLang/Reptile/engine"
	"regexp"
)

const cityListRe = `<a href="([^"]*)">([^<]+)</a>`
const Url = "http://www.supei.com/www/"

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	// 循环打印城市列表
	for _, m := range match {
		result.Items = append(result.Items, "City"+string(m[2]))
		result.Request = append(
			result.Request, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}
