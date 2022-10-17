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

	var limit = 5
	for _, m := range match {

		result.Items = append(result.Items, "City"+string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        Url + string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
