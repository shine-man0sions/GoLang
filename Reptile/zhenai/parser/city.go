package parser

import (
	"GoLang/Reptile/engine"
	"log"
	"regexp"
)

const cityRe = `<a href="(viewuser.jsp?[^"]+)"[^>]*>([^<]+)</a>`

// 城市解析器

func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}

	// 使用正则匹配匹配出名称和url
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	// 获取城市列表里的用户和用户链接
	for _, m := range matches {
		name := string(m[2])
		log.Printf("%s", name)
		result.Items = append(result.Items, "User"+name)
		result.Request = append(
			result.Request, engine.Request{
				Url: Url + string(m[1]),
				ParserFunc: func(contents []byte) engine.ParseResult {
					return ParseProfile(contents, name)
				},
			})
	}
	return result
}
