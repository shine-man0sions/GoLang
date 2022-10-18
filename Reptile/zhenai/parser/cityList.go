package parser

import (
	"GoLang/Reptile/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

// 城市列表解析器 城市名称+url

func ParseCityList(contents []byte) engine.ParseResult {

	// 使用正则匹配匹配出名称和url
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	// 循环打印城市列表
	for _, m := range matches {
		result.Items = append(result.Items, "City"+string(m[2]))
		result.Request = append(
			result.Request, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}
